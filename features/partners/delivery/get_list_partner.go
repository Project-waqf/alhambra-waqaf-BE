package delivery

import (
	"net/http"
	"strconv"
	"wakaf/features/partners/delivery/response"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (d PartnerDelivery) GetAllPartner() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		cnvLimit, err := strconv.Atoi(limit)
		if err != nil {
			d.logger.Error("Error convert offset", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Offset not valid"))
		}
		offset := c.QueryParam("offset")
		cnvOffset, err := strconv.Atoi(offset)
		if err != nil {
			d.logger.Error("Error convert offset", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Offset not valid"))
		}

		sort := c.QueryParam("sort")

		res, err := d.PartnerService.GetAllPartner(cnvLimit, cnvOffset, sort)
		if err != nil {
			d.logger.Error("Failed get list partner", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Getdetail partner successfully", response.GetListResponse(res)))
	}
}
