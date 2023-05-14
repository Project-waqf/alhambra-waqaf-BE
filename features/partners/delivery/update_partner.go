package delivery

import (
	"net/http"
	"strconv"
	"wakaf/features/partners/delivery/request"
	"wakaf/features/partners/delivery/response"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (d PartnerDelivery) UpdatePartner() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input request.PartnerRequest

		err := c.Bind(&input)
		if err != nil {
			d.logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		id := c.Param("id")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			d.logger.Error("Error get id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Id not valid"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileIdDb, err := d.PartnerService.GetFileId(cnvId)
			if err != nil {
				d.logger.Error("Failed to get fileId", zap.Error(err))
				return c.JSON(http.StatusNotFound, helper.Failed("Failed to get fileId"))
			}
			err = helper.Delete(fileIdDb)
			if err != nil {
				d.logger.Error("Failed delete image in imagekit", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
			}

			fileId, fileName, err := helper.Upload(c, file, fileheader, "partner")
			if err != nil {
				d.logger.Error("Error upload image", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
			input.FileId = fileId
			input.Picture = fileName
		}

		res, err := d.PartnerService.UpdatePartner(cnvId, request.ToDomainCreatePartner(input))
		if err != nil {
			d.logger.Error("Failed create partner", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Getdetail partner successfully", response.GetDetailResponse(res)))
	}
}
