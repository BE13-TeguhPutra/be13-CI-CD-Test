package factory

import (
	"be13/clean-arch/features/auth/delivery"
	"be13/clean-arch/features/auth/repository"
	"be13/clean-arch/features/auth/service"
	_delivery "be13/clean-arch/features/user/delivery"
	_repository "be13/clean-arch/features/user/repository"
	_service "be13/clean-arch/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(db *gorm.DB, e *echo.Echo) {
	userRepoFactroy := _repository.New(db)
	userServiceFactory := _service.New(userRepoFactroy)
	_delivery.New(userServiceFactory, e)

	authRepoFactory := repository.NewAuth(db)
	authServiceFactory := service.NewAuth(authRepoFactory)
	delivery.NewAuth(authServiceFactory, e)
}
