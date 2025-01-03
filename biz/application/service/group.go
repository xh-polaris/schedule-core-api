package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/consts"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/group"
)

type IGroupService interface {
	CreateDefaultGroup(ctx context.Context, userId string) (id string, err error)
	FindDefaultGroup(ctx context.Context, userId string) (id string, err error)
}
type GroupService struct {
	GroupMapper *group.MongoMapper
}

var GroupServiceSet = wire.NewSet(
	wire.Struct(new(GroupService), "*"),
	wire.Bind(new(IGroupService), new(*GroupService)),
)

func (s *GroupService) CreateDefaultGroup(ctx context.Context, userId string) (id string, err error) {
	g := &group.Group{
		Name:   "default",
		Status: 0,
	}
	id, err = s.GroupMapper.Insert(ctx, g)
	if err != nil {
		return "", consts.ErrDefaultGroup
	}
	return id, nil
}

func (s *GroupService) FindDefaultGroup(ctx context.Context, userId string) (id string, err error) {
	g, err := s.GroupMapper.FindDefault(ctx, userId)
	if err != nil {
		return "", err
	}
	return g.ID.Hex(), nil
}
