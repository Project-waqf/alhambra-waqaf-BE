package delivery

import (
	"net/http"
	"strings"
	"wakaf/features/admin/domain"
	"wakaf/middlewares"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AdminDelivery struct {
	AdminServices domain.UseCaseInterface
}

var (
	logger = helper.Logger()
)

func New(e *echo.Echo, data domain.UseCaseInterface) {

	handler := AdminDelivery{
		AdminServices: data,
	}

	e.POST("/admin/login", handler.Login())
	e.POST("/admin/register", handler.Register())
	e.PUT("/admin/update/password", handler.Edit(), middlewares.JWTMiddleware())
}

func (delivery *AdminDelivery) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Login
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		cnv := ToDomainLogin(input)

		res, err := delivery.AdminServices.Login(cnv)
		if err != nil {
			logger.Error("Login", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}

		return c.JSON(http.StatusOK, helper.Success("Login success", FromDomainLogin(res)))
	}
}

func (d *AdminDelivery) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Register
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		cnv := ToDomainRegister(input)

		err = d.AdminServices.Register(cnv)
		if err != nil {
			logger.Error("Register", zap.Any("Register Failed", err.Error()))
			if strings.Contains(err.Error(), "email has taken") {
				return c.JSON(http.StatusBadRequest, helper.Failed("Register failed"))
			}
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Register success", input))
	}
}

func (d *AdminDelivery) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input Register

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		id, _ := middlewares.DecodeToken(c)
		input.Id = id
		cnv := ToDomainRegister(input)
		err = d.AdminServices.UpdatePassword(cnv)
		if err != nil {
			logger.Error("Update", zap.Any("Update Password Failed", err.Error()))
			return c.JSON(http.StatusBadRequest, helper.Failed("Update Password Failed"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Update Password success", nil))
	}
}
