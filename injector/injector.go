//go:build wireinject
// +build wireinject

package injector

import (
	"projectcharter/controller"
	"projectcharter/repository"
	"projectcharter/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var projcharterSet = wire.NewSet(
	repository.NewProjectCharterRepository,
	service.NewProjectCharterService,
	controller.NewProjectCharterController,
)

func InitProjectCharter(db *gorm.DB) controller.ProjectCharterController {
	wire.Build(
		projcharterSet,
	)
	return nil
}
