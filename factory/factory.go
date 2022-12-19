package factory

import (
	"wakaf/features/user/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	handler.New(e)
}