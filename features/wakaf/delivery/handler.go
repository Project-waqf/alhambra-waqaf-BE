package delivery

import (
	"log"
	"net/http"
	"wakaf/config"
	"wakaf/features/wakaf/domain"
	"wakaf/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WakafDelivery struct {
	WakafService domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := WakafDelivery{
		WakafService: data,
	}

	e.POST("admin/wakaf", handler.AddWakaf(),  middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
}

func (wakaf *WakafDelivery) AddWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WakafRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		dest, err := helper.Upload(c, file, fileheader)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = dest
		res, err := wakaf.WakafService.AddWakaf(ToDomainAdd(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Somethin error in server"))
		}
		log.Print(res)
		return c.JSON(http.StatusOK, helper.Success("Add wakaf successfully", FromDomainAdd(res)))
	}
}
