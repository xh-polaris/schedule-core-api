package group

import (
	"context"
	"errors"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	prefixUserCacheKey = "cache:group"
	CollectionName     = "group"
)

type IMongoMapper interface {
	Insert(ctx context.Context, group Group) (id string, err error)
	Update(ctx context.Context, group Group) (err error)
	FindOne(ctx context.Context, id string) (group *Group, err error)
	FindDefault(ctx context.Context, userId string) (*Group, error)
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

func (m MongoMapper) Insert(ctx context.Context, group *Group) (id string, err error) {
	if group.ID.IsZero() {
		group.ID = primitive.NewObjectID()
		group.CreateTime = time.Now()
		group.UpdateTime = time.Now()
	}
	g, err := m.conn.InsertOneNoCache(ctx, group)
	if err != nil {
		return "", err
	}
	return g.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m MongoMapper) Update(ctx context.Context, group Group) (err error) {
	group.UpdateTime = time.Now()
	_, err = m.conn.UpdateByIDNoCache(ctx, group.ID, group)
	return err
}

func (m MongoMapper) FindOne(ctx context.Context, id string) (group *Group, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}
	var g Group
	err = m.conn.FindOneNoCache(ctx, &g, bson.M{
		consts.ID:     oid,
		consts.Status: consts.DefaultStatus,
	})
	if err != nil {
		return nil, consts.ErrNotFound
	}
	return &g, nil
}

func (m MongoMapper) FindDefault(ctx context.Context, userId string) (*Group, error) {
	var group Group
	err := m.conn.FindOneNoCache(ctx, &group, bson.M{
		consts.UserID: userId,
		consts.Name:   consts.DefaultGroupName,
		consts.Status: consts.DefaultStatus,
	})
	switch {
	case err == nil:
		return &group, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}
