package service

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/schedule-core-api/biz/adaptor"
	"github.com/xh-polaris/schedule-core-api/biz/application/dto/schedule/core_api"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/schedule"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IScheduleService interface {
	CreateSchedule(ctx context.Context, req *core_api.CreateScheduleReq) (*core_api.CreateScheduleResp, error)
	CreateScheduleFromOrigin(ctx context.Context, req *core_api.CreateScheduleFromOriginReq) (*core_api.CreateScheduleFromOriginResp, error)
	CreateSchedules(ctx context.Context, c *core_api.CreateSchedulesReq) (*core_api.CreateSchedulesResp, error)
	UpdateSchedule(ctx context.Context, req *core_api.UpdateScheduleReq) (*core_api.UpdateScheduleResp, error)
	GetSchedules(ctx context.Context, req *core_api.GetSchedulesReq) (*core_api.GetSchedulesResp, error)
}
type ScheduleService struct {
	ScheduleMapper *schedule.MongoMapper
	GroupService   *GroupService
}

var ScheduleServiceSet = wire.NewSet(
	wire.Struct(new(ScheduleService), "*"),
	wire.Bind(new(IScheduleService), new(*ScheduleService)),
)

// CreateSchedule 创建单个日程
func (s ScheduleService) CreateSchedule(ctx context.Context, req *core_api.CreateScheduleReq) (*core_api.CreateScheduleResp, error) {
	// 获取userId
	userId := adaptor.ExtractUserMeta(ctx).GetUserId()
	if userId == "" {
		return nil, consts.ErrNotAuthentication
	}

	groupId := req.Group
	if req.Group == "" {
		// 获取默认组id
		g, err := s.GroupService.FindDefaultGroup(ctx, userId)
		if err != nil {
			return nil, err
		}
		groupId = g
	}

	aSchedule := &schedule.Schedule{
		UserId:      userId,
		Group:       groupId,
		Title:       req.Title,
		Description: req.Description,
		Done:        consts.DefaultDone,
		Progress:    req.Progress,
		Priority:    req.Priority,
		Top:         req.Top,
		DDL:         time.Time{},
		Status:      consts.DefaultStatus,
	}

	// 0表示不设置ddl
	if req.Ddl != 0 {
		aSchedule.DDL = time.Unix(req.Ddl, 0)
	}

	err := s.ScheduleMapper.Insert(ctx, aSchedule)
	if err != nil {
		return nil, err
	}

	res := &core_api.Schedule{
		Id:          aSchedule.ID.Hex(),
		UserId:      aSchedule.UserId,
		Group:       aSchedule.Group,
		Origin:      aSchedule.Origin,
		Title:       aSchedule.Title,
		Description: aSchedule.Description,
		Done:        aSchedule.Done,
		Progress:    aSchedule.Progress,
		Priority:    aSchedule.Priority,
		Top:         aSchedule.Top,
		CreateTime:  aSchedule.CreateTime.Unix(),
		UpdateTime:  aSchedule.UpdateTime.Unix(),
	}

	// 零值就用0表示
	if !aSchedule.DDL.IsZero() {
		res.Ddl = aSchedule.DDL.Unix()
	} else {
		res.Ddl = 0
	}

	return &core_api.CreateScheduleResp{
		Code:     0,
		Msg:      "创建成功",
		Schedule: res,
	}, nil

}

// CreateScheduleFromOrigin 从原话生成日程，但是不入库
func (s ScheduleService) CreateScheduleFromOrigin(ctx context.Context, req *core_api.CreateScheduleFromOriginReq) (*core_api.CreateScheduleFromOriginResp, error) {
	// 获取userId
	userId := adaptor.ExtractUserMeta(ctx).GetUserId()
	if userId == "" {
		return nil, consts.ErrNotAuthentication
	}

	httpClient := util.NewHttpClient()
	response, err := httpClient.CallGLM(req.Origin)
	if err != nil {
		return nil, err
	}
	var simpleSchedulesJson string
	// 获取模型响应
	if simpleSchedulesJson, ok := response["choice"].([]map[string]interface{})[0]["message"].(map[string]interface{})["content"].(string); !ok || simpleSchedulesJson == "" {
		return nil, consts.ErrCall
	}
	var simpleSchedules []SimpleSchedule
	err = sonic.UnmarshalString(simpleSchedulesJson, &simpleSchedules)
	if err != nil {
		return nil, consts.ErrCall
	}
	var schedules []*core_api.Schedule
	for _, simpleSchedule := range simpleSchedules {
		aSchedule := &core_api.Schedule{
			Id:          primitive.NewObjectID().Hex(),
			UserId:      userId,
			Group:       "",
			Origin:      simpleSchedule.Origin,
			Title:       simpleSchedule.Title,
			Description: simpleSchedule.Description,
			Done:        consts.DefaultDone,
			Progress:    consts.DefaultProgress,
			Priority:    consts.DefaultPriority,
			Top:         consts.DefaultTop,
			Ddl:         0,
			CreateTime:  time.Now().Unix(),
			UpdateTime:  time.Now().Unix(),
		}
		ddl, err := time.Parse("0001-01-01 00:00:00", simpleSchedule.DDL)
		if err != nil {
			return nil, consts.ErrCall
		}
		if !ddl.IsZero() {
			aSchedule.Ddl = ddl.Unix()
		}
	}
	return &core_api.CreateScheduleFromOriginResp{
		Code:      0,
		Msg:       "success",
		Schedules: schedules,
	}, nil
}

// CreateSchedules 批量创建日程
func (s ScheduleService) CreateSchedules(ctx context.Context, req *core_api.CreateSchedulesReq) (*core_api.CreateSchedulesResp, error) {
	// 获取userId
	userId := adaptor.ExtractUserMeta(ctx).GetUserId()
	if userId == "" {
		return nil, consts.ErrNotAuthentication
	}

	defaultGroupId := ""

	// 应该用事务，插入失败就全部回滚，这里偷个懒
	for _, sc := range req.Schedules {
		oid, err := primitive.ObjectIDFromHex(sc.Id)
		if err != nil {
			return nil, consts.ErrInvalidObjectId
		}

		groupId := sc.Group
		if sc.Group == "" && defaultGroupId == "" {
			// 获取默认组id
			defaultGroupId, err = s.GroupService.FindDefaultGroup(ctx, userId)
			if err != nil {
				return nil, err
			}
			groupId = defaultGroupId
		} else if sc.Group == "" {
			groupId = defaultGroupId
		}

		aSchedule := &schedule.Schedule{
			ID:          oid,
			UserId:      sc.UserId,
			Group:       groupId,
			Origin:      sc.Origin,
			Title:       sc.Title,
			Description: sc.Description,
			Done:        sc.Done,
			Progress:    sc.Progress,
			Priority:    sc.Priority,
			Top:         sc.Top,
			DDL:         time.Time{},
			Status:      consts.DefaultStatus,
			CreateTime:  time.Unix(sc.CreateTime, 0),
			UpdateTime:  time.Unix(sc.UpdateTime, 0),
		}
		if sc.Ddl != 0 {
			aSchedule.DDL = time.Unix(sc.Ddl, 0)
		}
		err = s.ScheduleMapper.Insert(ctx, aSchedule)
		if err != nil {
			return nil, consts.ErrInsert
		}
	}

	res := &core_api.CreateSchedulesResp{
		Code: 0,
		Msg:  "success",
	}

	return res, nil
}

func (s ScheduleService) UpdateSchedule(ctx context.Context, req *core_api.UpdateScheduleReq) (*core_api.UpdateScheduleResp, error) {
	// 获取userId
	userId := adaptor.ExtractUserMeta(ctx).GetUserId()
	if userId == "" {
		return nil, consts.ErrNotAuthentication
	}

	aSchedule := req.Schedule
	oid, err := primitive.ObjectIDFromHex(aSchedule.Id)
	if err != nil {
		return nil, consts.ErrRepeatedSignUp
	}

	if userId != aSchedule.UserId {
		return nil, consts.ErrForbidden
	}

	newSchedule := &schedule.Schedule{
		ID:          oid,
		UserId:      aSchedule.UserId,
		Group:       aSchedule.Group,
		Origin:      aSchedule.Origin,
		Title:       aSchedule.Title,
		Description: aSchedule.Description,
		Done:        aSchedule.Done,
		Progress:    aSchedule.Progress,
		Priority:    aSchedule.Priority,
		Top:         aSchedule.Top,
		DDL:         time.Time{},
	}
	// 更新ddl
	if aSchedule.Ddl != 0 {
		newSchedule.DDL = time.Unix(aSchedule.Ddl, 0)
	}

	err = s.ScheduleMapper.Update(ctx, newSchedule)
	if err != nil {
		return nil, err
	}
	res := &core_api.Schedule{
		Id:          newSchedule.ID.Hex(),
		UserId:      newSchedule.UserId,
		Group:       newSchedule.Group,
		Origin:      newSchedule.Origin,
		Title:       newSchedule.Title,
		Description: newSchedule.Description,
		Done:        newSchedule.Done,
		Progress:    newSchedule.Progress,
		Priority:    newSchedule.Priority,
		Top:         newSchedule.Top,
		CreateTime:  newSchedule.CreateTime.Unix(),
		UpdateTime:  newSchedule.UpdateTime.Unix(),
	}

	// 零值就用0表示
	if !newSchedule.DDL.IsZero() {
		res.Ddl = newSchedule.DDL.Unix()
	} else {
		res.Ddl = 0
	}

	return &core_api.UpdateScheduleResp{
		Code: 0,
		Msg:  "创建成功",
	}, nil
}

func (s ScheduleService) GetSchedules(ctx context.Context, req *core_api.GetSchedulesReq) (*core_api.GetSchedulesResp, error) {
	// 获取userId
	userId := adaptor.ExtractUserMeta(ctx).GetUserId()
	if userId == "" {
		return nil, consts.ErrNotAuthentication
	}
	data, total, err := s.ScheduleMapper.FindMany(ctx, userId, req.PaginationOptions)
	if err != nil {
		return nil, err
	}
	var schedules []*core_api.Schedule
	for _, v := range data {
		aSchedule := &core_api.Schedule{}
		err = copier.Copy(aSchedule, v)
		if err != nil {
			return nil, err
		}
		aSchedule.Id = v.ID.Hex()
		if v.DDL.IsZero() {
			aSchedule.Ddl = 0
		} else {
			aSchedule.Ddl = v.DDL.Unix()
		}
		aSchedule.CreateTime = v.CreateTime.Unix()
		aSchedule.UpdateTime = v.UpdateTime.Unix()
		schedules = append(schedules, aSchedule)
	}
	return &core_api.GetSchedulesResp{
		Total:     total,
		Schedules: schedules,
	}, err
}

type SimpleSchedule struct {
	Origin      string `json:"origin"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DDL         string `json:"ddl"`
}
