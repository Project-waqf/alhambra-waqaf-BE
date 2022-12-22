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

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		
		filename, err := helper.Upload(c, file, fileheader)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		
		input.Picture = filename
		cnv := ToDomain(input)
		res, err := news.NewsServices.AddNews(cnv)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}
		addResponse := FromDomain(res)
		return c.JSON(http.StatusBadRequest, helper.Success("Add News Successfully", addResponse))
	}
}