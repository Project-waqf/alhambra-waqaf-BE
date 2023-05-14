package delivery

import (
	"net/http"
	"wakaf/features/partners/delivery/request"
	"wakaf/features/partners/delivery/response"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (d PartnerDelivery) CreatePartner() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input request.PartnerRequest

		err := c.Bind(&input)
		if err != nil {
			d.logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			d.logger.Error("Error get picture", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		fileId, dest, err := helper.Upload(c, file, fileheader, "partner")
		if err != nil {
			d.logger.Error("Error upload images", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = dest
		input.FileId = fileId
		res, err := d.PartnerService.CreatePartner(request.ToDomainCreatePartner(input))
		if err != nil {
			d.logger.Error("Failed create partner", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Add wakaf successfully", response.GetDetailResponse(res)))
	}
}
