package handler

import (
	"log"
	"net/http"
	"wakaf/helper"

	"github.com/labstack/echo/v4"
)

type AdminDelivery struct{

}

func New(e *echo.Echo) {

	handler := AdminDelivery{}

	e.POST("/admin/users", handler.Upload())

}

func (upload *AdminDelivery) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		
		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			err := helper.Upload(c, file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusOK, "Error")
			}
		}
		return c.JSON(http.StatusOK, "Sukses")
	}
}