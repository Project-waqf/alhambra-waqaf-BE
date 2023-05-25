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
	e.POST("/admin/forgot", handler.Forgot())
	e.PUT("/admin/profile", handler.Edit(), middlewares.JWTMiddleware())
	e.POST("/admin/forgot/update", handler.UpdateForgot())
	e.PUT("/admin/profile/image", handler.UpdateImage(), middlewares.JWTMiddleware())
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
			if strings.Contains(err.Error(), "not found") {
				ret := echo.ErrNotFound
				ret.Message = "email not found"
				ret.Code = http.StatusNotFound
				ret.Internal = err
				return ret
			} else if strings.Contains(err.Error(), "password invalid") {
				ret := echo.ErrUnauthorized
				ret.Message = "wrong password"
				return ret
			}
			return echo.ErrBadRequest
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
		return c.JSON(http.StatusCreated, helper.Success("Register success", RegisterResponseNew{Name: input.Name, Email: input.Email, Password: input.Password}))
	}
}

func (d *AdminDelivery) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = Register{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		if input.OldPassword != "" {
			_, err = d.AdminServices.Login(ToDomainLogin(Login{Email: input.Email, Password: input.OldPassword}))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.Failed("wrong password"))
			}
		}

		id, _ := middlewares.DecodeToken(c)
		input.Id = id
		cnv := ToDomainRegister(input)
		res, err := d.AdminServices.UpdateProfile(cnv)
		if err != nil {
			logger.Error("Update profile failed", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Update Password Failed"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Update Password success", FromDomainProfile(res)))
	}
}

func (d *AdminDelivery) Forgot() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Forgot

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		res, err := d.AdminServices.ForgotSendEmail(domain.Admin{Email: input.Email})
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.Failed("Email not found"))
			}
			return c.JSON(http.StatusBadRequest, helper.Failed("Send Email Failed"))
		}
		return c.JSON(http.StatusOK, helper.Success("Send Email success", FromDomainLogin(res)))
	}
}

func (d *AdminDelivery) UpdateForgot() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ForgotUpdate
		err := c.Bind(&input)
		if err != nil {
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		err = d.AdminServices.ForgotUpdate(input.Token, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "token not valid") {
				return c.JSON(http.StatusBadRequest, helper.Failed("Token not valid"))
			}
			return c.JSON(http.StatusBadRequest, helper.Failed("Reset Password Failed"))
		}

		return c.JSON(http.StatusCreated, helper.Success("Reset Password success", nil))
	}
}

func (d *AdminDelivery) UpdateImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middlewares.DecodeToken(c)

		resProfile, err := d.AdminServices.GetProfile(uint(id))
		if err != nil {
			logger.Error("Error get profile admin", zap.Error(err))
			ret := echo.ErrInternalServerError
			ret.Message = "Error get profile admin"
			return ret
		}

		// Delete image in imagekit if exist
		if resProfile.Image != "" {
			if err := helper.Delete(resProfile.FileId); err != nil {
				return echo.ErrInternalServerError
			}
		}
		
		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			logger.Error("Error get picture", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		fileId, dest, err := helper.Upload(c, file, fileheader, "profile")
		if err != nil {
			logger.Error("Error upload images", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		
		var input = domain.Admin{
			ID: uint(id),
			Image: dest,
			FileId: fileId,
		}

		err = d.AdminServices.UpdateImage(input)
		if err != nil {
			ret := echo.ErrInternalServerError
			ret.Message = "Failed update profile image"
			return ret
		}
		return c.JSON(http.StatusCreated, helper.Success("Success update profile images", dest))
	}
}