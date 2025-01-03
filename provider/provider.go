package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/schedule-core-api/biz/application/service"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/group"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/schedule"
	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/mapper/user"

	"github.com/xh-polaris/schedule-core-api/biz/infrastructure/config"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config          *config.Config
	UserService     service.UserService
	ScheduleService service.ScheduleService
	GroupService    service.GroupService
}

func Get() *Provider {
	return provider
}

var ApplicationSet = wire.NewSet(
	service.UserServiceSet,
	service.ScheduleServiceSet,
	service.GroupServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	user.NewMongoMapper,
	schedule.NewMongoMapper,
	group.NewMongoMapper,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	InfrastructureSet,
)
