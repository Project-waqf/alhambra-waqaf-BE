package delivery

import (
	"net/http"
	"wakaf/features/news/domain"
	"wakaf/helper"

	"github.com/labstack/echo/v4"
)

type NewsDelivery struct {
	NewsServices domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := NewsDelivery{
		NewsServices: data,
	}

	e.POST("/admin/news", handler.AddNews())
}

func (news *NewsDelivery) AddNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input News
		
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		cnv := ToDomain(input)
		res, err := news.NewsServices.AddNews(cnv)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}
		addResponse := FromDomain(res)
		return c.JSON(http.StatusBadRequest, helper.Success("Add News Successfully", addResponse))
	}
}