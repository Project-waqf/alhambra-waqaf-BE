package delivery

import (
	"log"
	"net/http"
	"wakaf/config"
	"wakaf/features/asset/domain"
	"wakaf/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AssetDelivery struct {
	AssetService domain.UsecaseInterface
}

func New(e *echo.Echo, data domain.UsecaseInterface) {
	handler := &AssetDelivery{
		AssetService: data,
	}

	e.POST("admin/asset", handler.AddAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
}

func (asset *AssetDelivery) AddAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AssetRequest

		err := c.Bind(&input)
		if err != nil {
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
			}
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		fileName, err := helper.Upload(c, file, fileheader)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = fileName
		res, err := asset.AssetService.AddAsset(ToDomainAdd(input))
		if err != nil {
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
		}
		return c.JSON(http.StatusOK, helper.Success("Add asset successfully", res))
	}
}