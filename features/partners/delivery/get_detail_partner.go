package delivery

import (
	"net/http"
	"strconv"
	"wakaf/features/partners/delivery/response"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (d PartnerDelivery) GetDetailPartner() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			d.logger.Error("Error get id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Id not valid"))
		}

		res, err := d.PartnerService.GetPartnerDetail(cnvId)
		if err != nil {
			d.logger.Error("Failed get detail partner", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Get detail partner successfully", response.GetDetailResponse(res)))
	}
}
