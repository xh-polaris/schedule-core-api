package schedule

import (
	"context"
	"github.com/xh-polaris/schedule-core-api/biz/application/dto/basic"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/consts"
	util "github.com/xh-polaris/schedule-core-api/biz/infrastructure/util/page"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	prefixSchedulerCacheKey = "cache:schedule"
	CollectionName          = "schedule"
)

type IMongoMapper interface {
	Insert(ctx context.Context, schedule *Schedule) error
	Update(ctx context.Context, schedule *Schedule) error
	FindOne(ctx context.Context, id string) (*Schedule, error)
	FindMany(ctx context.Context, userId string, p *basic.PaginationOptions) (data []*Schedule, total int64, err error)
	DeleteOne(ctx context.Context, userId string, id string) error
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{
		conn: conn,
	}
}

func (m *MongoMapper) Insert(ctx context.Context, schedule *Schedule) error {
	if schedule.ID.IsZero() {
		schedule.ID = primitive.NewObjectID()
		schedule.CreateTime = time.Now()
		schedule.UpdateTime = schedule.CreateTime
	}
	_, err := m.conn.InsertOneNoCache(ctx, schedule)
	return err
}

func (m *MongoMapper) Update(ctx context.Context, schedule *Schedule) error {
	schedule.UpdateTime = time.Now()
	_, err := m.conn.UpdateByIDNoCache(ctx, schedule.ID, bson.M{consts.SET: schedule})
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (*Schedule, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}
	var u Schedule
	err = m.conn.FindOneNoCache(ctx, &u, bson.M{
		consts.ID:     oid,
		consts.Status: consts.DefaultStatus,
	})
	if err != nil {
		return nil, consts.ErrNotFound
	}
	return &u, nil
}

func (m *MongoMapper) FindMany(ctx context.Context, userId string, p *basic.PaginationOptions) (data []*Schedule, total int64, err error) {
	skip, limit := util.ParsePageOpt(p)
	data = make([]*Schedule, 0, limit)
	err = m.conn.Find(ctx, &data,
		bson.M{
			consts.UserID: userId,
			consts.Status: consts.DefaultStatus,
		}, &options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort: bson.D{
				{Key: consts.Top, Value: -1},
				{Key: consts.Priority, Value: -1},
				{Key: consts.CreateTime, Value: -1},
			},
		})
	if err != nil {
		return nil, 0, err
	}

	total, err = m.conn.CountDocuments(ctx, bson.M{
		consts.UserID: userId,
		consts.Status: consts.DefaultStatus,
	})
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (m *MongoMapper) DeleteOne(ctx context.Context, userId string, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInvalidObjectId
	}
	_, err = m.conn.UpdateOneNoCache(ctx, bson.M{
		consts.ID:     oid,
		consts.UserID: userId,
	}, bson.M{
		"$set": bson.M{
			consts.Status:     consts.DefaultStatus,
			consts.DeleteTime: time.Now(),
		},
	})
	return err
}
