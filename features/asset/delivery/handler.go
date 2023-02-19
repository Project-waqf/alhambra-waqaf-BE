package delivery

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"wakaf/config"
	"wakaf/features/asset/domain"
	"wakaf/pkg/helper"

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

	e.POST("/admin/asset", handler.AddAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.GET("/admin/asset", handler.GetAllAsset())
	e.GET("/admin/asset/:id_asset", handler.GetAsset())
	e.PUT("/admin/asset/:id_asset", handler.UpdateAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.DELETE("/admin/asset/:id_asset", handler.DeleteAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.PUT("/admin/asset/online/:id_asset", handler.ToOnline(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
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

		fileId, fileName, err := helper.Upload(c, file, fileheader, "asset")
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = fileName
		input.FileId = fileId
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
			fileId, fileName, err := helper.Upload(c, file, fileheader, "asset")
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
			}
			input.FileId = fileId
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

func (asset *AssetDelivery) DeleteAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_asset")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		err = asset.AssetService.DeleteAsset(uint(cnvId))
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, helper.Failed("Data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
		}
		return c.JSON(http.StatusOK, helper.Success("Delete asset successfully", nil))
	}
}

func (asset *AssetDelivery) ToOnline() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_asset")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		err = asset.AssetService.ToOnline(uint(cnvId))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Set asset to online successfully", nil))
	}
}