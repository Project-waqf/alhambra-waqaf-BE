package delivery

import (
	"net/http"
	"strconv"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (d PartnerDelivery) DeletePartner() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			d.logger.Error("Error get id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Id not valid"))
		}

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

		err = d.PartnerService.DeletePartner(cnvId)
		if err != nil {
			d.logger.Error("Failed delete partner", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Delete partner successfully", nil))

	}
}
