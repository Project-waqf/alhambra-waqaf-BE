package factory

import (
	AdminDelivery "wakaf/features/admin/delivery"
	AdminRepository "wakaf/features/admin/repository"
	AdminServices "wakaf/features/admin/usecase"
	AssetDelivery "wakaf/features/asset/delivery"
	AssetRepository "wakaf/features/asset/repository"
	AssetServices "wakaf/features/asset/services"
	NewsDelivery "wakaf/features/news/delivery"
	NewsRepository "wakaf/features/news/repository"
	NewsServices "wakaf/features/news/services"
	WakafDelivery "wakaf/features/wakaf/delivery"
	WakafRepository "wakaf/features/wakaf/repository"
	WakafServices "wakaf/features/wakaf/services"
	PartnerDelivery "wakaf/features/partners/delivery"
	PartnerRepository "wakaf/features/partners/repository"
	ParnterServices "wakaf/features/partners/services"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, redis *redis.Client, logger *zap.Logger) {
	// ADMIN
	adminRepoFactory := AdminRepository.New(db, redis)
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

	// PARTNER
	partnerRepoFactory := PartnerRepository.New(db)
	partnerServiceFactory := ParnterServices.New(partnerRepoFactory, logger)
	PartnerDelivery.New(e, partnerServiceFactory, logger)
}
