package delivery

import (
	"log"
	"net/http"
	"strconv"
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
	e.GET("admin/asset", handler.GetAllAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.GET("admin/asset/:id_asset", handler.GetAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.PUT("admin/asset/:id_asset", handler.UpdateAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))

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
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Add asset successfully", FromDomainAdd(res)))
	}
}

func (asset *AssetDelivery) GetAllAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := asset.AssetService.GetAllAsset()
		if err != nil {
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
		}
		return c.JSON(http.StatusOK, helper.Success("Add asset successfully", FromDomainGetAll(res)))
	}
}

func (asset *AssetDelivery) GetAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_asset")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		res, err := asset.AssetService.GetAsset(uint(cnvId))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Get asset successfully", FromDomainAdd(res)))
	}
}

func (asset *AssetDelivery) UpdateAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AssetRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileName, err := helper.Upload(c, file, fileheader)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
			}
			input.Picture = fileName
		}

		id := c.Param("id_asset")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		res, err := asset.AssetService.UpdateAsset(uint(cnvId), ToDomainAdd(input))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Update asset successfully", FromDomainAdd(res)))
	}
}