package delivery

import (
	"wakaf/config"
	"wakaf/features/partners/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type PartnerDelivery struct {
	PartnerService domain.UseCaseInterface
	logger *zap.Logger
}

func New(e *echo.Echo, data domain.UseCaseInterface, logger *zap.Logger) {
	handler := PartnerDelivery{
		PartnerService: data,
		logger: logger,
	}

	e.GET("/partners", handler.GetAllPartner())
	e.GET("/partner/:id", handler.GetDetailPartner())
	e.POST("/partner", handler.CreatePartner(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.PUT("/partner/:id", handler.UpdatePartner(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.DELETE("/partner/:id", handler.DeletePartner(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
}
