package factory

import (
	AdminRepository "wakaf/features/admin/repository"
	AdminServices "wakaf/features/admin/services"
	AdminDelivery "wakaf/features/admin/delivery"
	NewsRepository "wakaf/features/news/repository"
	NewsServices "wakaf/features/news/services"
	NewsDelivery "wakaf/features/news/delivery"
	WakafRepository "wakaf/features/wakaf/repository"
	WakafServices "wakaf/features/wakaf/services"
	WakafDelivery "wakaf/features/wakaf/delivery"
	AssetRepository "wakaf/features/asset/repository"
	AssetServices "wakaf/features/asset/services"
	AssetDelivery "wakaf/features/asset/delivery"
	

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	// ADMIN
	adminRepoFactory := AdminRepository.New(db)
	adminServiceFactory := AdminServices.New(adminRepoFactory)
	AdminDelivery.New(e, adminServiceFactory)

	// NEWS
	newsRepoFactory := NewsRepository.New(db)
	newsServiceFactory := NewsServices.New(newsRepoFactory)
	NewsDelivery.New(e, newsServiceFactory)

	// WAKAF
	wakafRepoFactory := WakafRepository.New(db)
	wakafServiceFactory := WakafServices.New(wakafRepoFactory)
	WakafDelivery.New(e, wakafServiceFactory)

	// ASSET
	assetRepoFactory := AssetRepository.New(db)
	assetServiceFactory := AssetServices.New(assetRepoFactory)
	AssetDelivery.New(e, assetServiceFactory)
}