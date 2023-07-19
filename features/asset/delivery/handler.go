package delivery

import (
	"net/http"
	"strconv"
	"strings"
	"wakaf/config"
	"wakaf/features/asset/domain"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type AssetDelivery struct {
	AssetService domain.UsecaseInterface
}

var (
	logger = helper.Logger()
)

func New(e *echo.Echo, data domain.UsecaseInterface) {
	handler := &AssetDelivery{
		AssetService: data,
	}

	e.POST("/admin/asset", handler.AddAsset(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.GET("/asset", handler.GetAllAsset())
	e.GET("/asset/:id_asset", handler.GetAsset())
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
				logger.Error("Error bind data", zap.Error(err))
				return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
			}
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			logger.Error("Error get image", zap.Error(err))
			fileId, fileName, err := helper.Upload(c, file, fileheader, "asset")
			if err != nil {
				logger.Error("Failed upload image", zap.Error(err))
				fileName = ""
				fileId = ""
			}
			input.Picture = fileName
			input.FileId = fileId
		}

		res, err := asset.AssetService.AddAsset(ToDomainAdd(input))
		if err != nil {
			logger.Error("Error in usecase", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Add asset successfully", FromDomainAdd(res)))
	}
}

func (asset *AssetDelivery) GetAllAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		status := c.QueryParam("status")
		page := c.QueryParam("page")
		sort := c.QueryParam("sort")

		cnvPage, err := strconv.Atoi(page)
		if err != nil {
			logger.Error("Failed to convert query param page")
		}
		res, countOnline, countDraft, countArchive, err := asset.AssetService.GetAllAsset(status, cnvPage, sort)
		if err != nil {
			if err != nil {
				logger.Error("Error in usecase", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
		}
		return c.JSON(http.StatusOK, helper.SuccessGetAll("Get all asset successfully", FromDomainGetAll(res), countOnline, countDraft, countArchive))
	}
}

func (asset *AssetDelivery) GetAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_asset")

		cnvId, err := strconv.Atoi(id)
		if err != nil {
			logger.Error("Error convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		res, err := asset.AssetService.GetAsset(uint(cnvId))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				logger.Error("Data not found", zap.Error(err))
				return c.JSON(http.StatusNotFound, helper.Failed("Asset not found"))
			}
			logger.Error("Error in usecase", zap.Error(err))
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
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		id := c.Param("id_asset")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileIdDb, err := asset.AssetService.GetFileId(uint(cnvId))
			if err != nil {
				logger.Error("Failed to get fileId", zap.Error(err))
				return c.JSON(http.StatusNotFound, helper.Failed("Id not found"))
			} else if err == nil && fileIdDb == "" {
				fileId, fileName, err := helper.Upload(c, file, fileheader, "asset")
				if err != nil {
					logger.Error("Failed upload image to imagekit", zap.Error(err))
					return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
				}
				input.FileId = fileId
				input.Picture = fileName
			} else {
				err = helper.Delete(fileIdDb)
				if err != nil {
					logger.Error("Failed delete image in imagekit", zap.Error(err))
					return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
				}
				fileId, fileName, err := helper.Upload(c, file, fileheader, "asset")
				if err != nil {
					logger.Error("Failed upload image to imagekit", zap.Error(err))
					return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
				}
				input.FileId = fileId
				input.Picture = fileName
			}
		}

		res, err := asset.AssetService.UpdateAsset(uint(cnvId), ToDomainAdd(input))
		if err != nil {
			logger.Error("Error in usecase", zap.Error(err))
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
			logger.Error("Error convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		fileIdDb, err := asset.AssetService.GetFileId(uint(cnvId))
		if err == nil && fileIdDb != "" {
			err = helper.Delete(fileIdDb)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
			}
		}

		err = asset.AssetService.DeleteAsset(uint(cnvId))
		if err != nil {
			logger.Error("Error in usecase", zap.Error(err))
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
			logger.Error("Error convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		err = asset.AssetService.ToOnline(uint(cnvId))
		if err != nil {
			logger.Error("Failed change asset to online", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Set asset to online successfully", nil))
	}
}
