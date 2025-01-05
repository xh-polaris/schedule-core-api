package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/util"

	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/user"

	"github.com/xh-polaris/schedule-core-api/biz/application/dto/schedule/core_api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IUserService interface {
	SignUp(ctx context.Context, req *core_api.SignUpReq) (*core_api.SignUpResp, error)
	SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error)
}
type UserService struct {
	UserMapper   *user.MongoMapper
	GroupService *GroupService
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (u *UserService) SignUp(ctx context.Context, req *core_api.SignUpReq) (*core_api.SignUpResp, error) {
	// 查找数据库判断手机号是否注册过
	oldUser, err := u.UserMapper.FindOneByPhone(ctx, req.AuthId)
	if err == nil && oldUser != nil {
		return nil, consts.ErrRepeatedSignUp
	} else if err != nil && !errors.Is(err, consts.ErrNotFound) {
		return nil, consts.ErrSignUp
	}

	// 在中台注册账户
	httpClient := util.NewHttpClient()
	signUpResponse, err := httpClient.SignUp(req.AuthType, req.AuthId, &req.VerifyCode)
	if err != nil {
		return nil, consts.ErrSignUp
	}
	userId := signUpResponse["userId"].(string)

	// 在中台设置密码
	authorization := signUpResponse["accessToken"].(string)
	if req.Password != "" {
		_, err = httpClient.SetPassword(authorization, req.Password)
		if err != nil {
			return nil, consts.ErrSignUp
		}
	}

	// 初始化用户
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	aUser := user.User{
		ID:         oid,
		Phone:      req.AuthId,
		Status:     0,
		CreateTime: now,
		UpdateTime: now,
	}

	// 向数据库中插入数据
	err = u.UserMapper.Insert(ctx, &aUser)
	if err != nil {
		return nil, consts.ErrSignUp
	}
	_, err = u.GroupService.CreateDefaultGroup(ctx, aUser.ID.Hex())
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &core_api.SignUpResp{
		Id:           userId,
		AccessToken:  authorization,
		AccessExpire: int64(signUpResponse["accessExpire"].(float64)),
	}, nil
}

func (u *UserService) SignIn(ctx context.Context, req *core_api.SignInReq) (*core_api.SignInResp, error) {
	// 查找数据库判断手机号是否注册过
	aUser, err := u.UserMapper.FindOneByPhone(ctx, req.AuthId)
	if errors.Is(err, consts.ErrNotFound) || aUser == nil { // 未找到，说明没有注册
		return nil, consts.ErrNotSignUp
	} else if err != nil {
		return nil, consts.ErrSignUp
	}

	// 通过中台登录
	httpClient := util.NewHttpClient()
	signInResponse, err := httpClient.SignIn(req.AuthType, req.AuthId, req.VerifyCode, req.Password)
	if err != nil {
		return nil, consts.ErrSignIn
	}
	userId := signInResponse["userId"].(string)

	return &core_api.SignInResp{
		Id:           userId,
		AccessToken:  signInResponse["accessToken"].(string),
		AccessExpire: int64(signInResponse["accessExpire"].(float64)),
	}, nil
}

func (s *UserService) SendVerifyCode(ctx context.Context, req *core_api.SendVerifyCodeReq) (*core_api.SendVerifyCodeResp, error) {
	aUser, err := s.UserMapper.FindOneByPhone(ctx, req.AuthId)
	if req.Type == 1 { // 登录验证码
		// 查找数据库判断手机号是否注册过
		if errors.Is(err, consts.ErrNotFound) || aUser == nil { // 未找到，说明没有注册
			return nil, consts.ErrNotSignUp
		} else if err != nil {
			return nil, consts.ErrSend
		}
	} else {
		if err == nil && aUser != nil {
			return nil, consts.ErrRepeatedSignUp
		} else if err != nil && !errors.Is(err, consts.ErrNotFound) {
			return nil, consts.ErrSignUp
		}
	}

	// 通过中台发送验证码
	httpClient := util.NewHttpClient()
	_, err = httpClient.SendVerifyCode(req.AuthType, req.AuthId)
	if err != nil {
		return nil, consts.ErrSend
	}

	return &core_api.SendVerifyCodeResp{
		Code: 0,
		Msg:  "发送验证码成功，请注意查收",
	}, nil
}

//func (u *UserService) GetUserInfo(ctx context.Context, req *core_api.GetUserInfoReq) (*core_api.GetUserInfoResp, error) {
//	userMeta := adaptor.ExtractUserMeta(ctx)
//	if userMeta.GetUserId() == "" {
//		return nil, consts.ErrNotAuthentication
//	}
//	aUser, err := u.UserMapper.FindOne(ctx, userMeta.GetUserId())
//	if err != nil {
//		return &core_api.GetUserInfoResp{
//			Code:    -1,
//			Msg:     "查询用户信息失败，请先登录或重试",
//			Payload: nil,
//		}, nil
//	}
//	return &core_api.GetUserInfoResp{
//		Code: 0,
//		Msg:  "查询成功",
//		Payload: &core_api.GetUserInfoResp_Payload{
//			Name:  aUser.Username,
//			Count: aUser.Count,
//			Phone: aUser.Phone,
//		},
//	}, nil
//}
//
//func (u *UserService) UpdateUserInfo(ctx context.Context, req *core_api.UpdateUserInfoReq) (*core_api.Response, error) {
//	// 获取用户id
//	userMeta := adaptor.ExtractUserMeta(ctx)
//	if userMeta.GetUserId() == "" {
//		return nil, consts.ErrNotAuthentication
//	}
//
//	// 根据用户id查询这个用户
//	aUser, err := u.UserMapper.FindOne(ctx, userMeta.GetUserId())
//	if err != nil {
//		return nil, consts.ErrNotFound
//	}
//
//	// 更新用户信息
//	aUser.Username = req.Name
//
//	// 存入新的用户信息
//	err = u.UserMapper.Update(ctx, aUser)
//	if err != nil {
//		return nil, consts.ErrUpdate
//	}
//
//	// 返回响应
//	return &core_api.Response{
//		Code: 0,
//		Msg:  "更新成功",
//	}, nil
//}
//
//func (u *UserService) UpdatePassword(ctx context.Context, req *core_api.UpdatePasswordReq) (*core_api.UpdatePasswordResp, error) {
//	// 获取用户id
//	userMeta := adaptor.ExtractUserMeta(ctx)
//	if userMeta.GetUserId() == "" {
//		return nil, consts.ErrNotAuthentication
//	}
//
//	// 根据用户id查询这个用户
//	aUser, err := u.UserMapper.FindOne(ctx, userMeta.GetUserId())
//	if err != nil {
//		return nil, consts.ErrNotFound
//	}
//
//	// 在中台注册账户
//	httpClient := util.NewHttpClient()
//	signInResponse, err := httpClient.SignUp(consts.Phone, aUser.Phone, &req.VerifyCode)
//	if err != nil {
//		return nil, consts.ErrVerifyCode
//	}
//
//	// 在中台设置密码
//	authorization := signInResponse["accessToken"].(string)
//	_, err = httpClient.SetPassword(authorization, req.Password)
//	if err != nil {
//		return nil, consts.ErrSignUp
//	}
//	return &core_api.UpdatePasswordResp{
//		Id:           aUser.ID.Hex(),
//		AccessToken:  authorization,
//		AccessExpire: int64(signInResponse["accessExpire"].(float64)),
//		Name:         aUser.Username,
//	}, nil
//}
