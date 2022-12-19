package factory

import (
	AdminRepository "wakaf/features/admin/repository"
	AdminServices "wakaf/features/admin/services"
	AdminDelivery "wakaf/features/admin/delivery"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	adminRepoFactory := AdminRepository.New(db)
	adminServiceFactory := AdminServices.New(adminRepoFactory)
	AdminDelivery.New(e, adminServiceFactory)
	
}