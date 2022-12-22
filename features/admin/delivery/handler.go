package delivery

import (
	"log"
	"net/http"
	"wakaf/features/admin/domain"
	"wakaf/helper"
	"wakaf/middlewares"
	"github.com/labstack/echo/v4"
)

type AdminDelivery struct{
	AdminServices domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {

	handler := AdminDelivery{
		AdminServices: data,
	}

	e.POST("/admin/login", handler.Login())

}

func (delivery *AdminDelivery) Login() echo.HandlerFunc{
	return func(c echo.Context) error {
		var input Login
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		cnv := ToDomainLogin(input)

		res, err := delivery.AdminServices.Login(cnv)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}
		loginRes := FromDomainLogin(res)
		loginRes.Token, _= middlewares.CreateToken(int(res.ID), res.Username)
		return c.JSON(http.StatusBadRequest, helper.Success("Login success", loginRes))
	}
}