package delivery

import (
	"log"
	"net/http"
	"strconv"
	"wakaf/config"
	"wakaf/features/wakaf/domain"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type WakafDelivery struct {
	WakafService domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := WakafDelivery{
		WakafService: data,
	}

	e.POST("/admin/wakaf", handler.AddWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))                 // INSERT WAKAF
	e.GET("/admin/wakaf", handler.GetAllWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))               // GET ALL WAKAF
	e.PUT("/admin/wakaf/:id_wakaf", handler.UpdateWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))     // UPDATE WAKAF
	e.DELETE("/admin/wakaf/:id_wakaf", handler.DeleteWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT))) // DELETE WAKAF
}

func (wakaf *WakafDelivery) AddWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WakafRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		dest, err := helper.Upload(c, file, fileheader)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = dest
		res, err := wakaf.WakafService.AddWakaf(ToDomainAdd(input))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Add wakaf successfully", FromDomainAdd(res)))
	}
}

func (wakaf *WakafDelivery) GetAllWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := wakaf.WakafService.GetAllWakaf()
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Get all wakaf successfully", FromDomainGetAll(res)))
	}
}

func (wakaf *WakafDelivery) UpdateWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WakafRequest

		id := c.Param("id_wakaf")
		err := c.Bind(&input)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileName, err := helper.Upload(c, file, fileheader)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
			input.Picture = fileName
		}

		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		res, err := wakaf.WakafService.UpdateWakaf(uint(cnvId), ToDomainAdd(input))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Update wakaf successfully", FromDomainAdd(res)))
	}
}

func (wakaf *WakafDelivery) DeleteWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_wakaf")

		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		_, err = wakaf.WakafService.DeleteWakaf(uint(cnvId))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Delete wakaf successfully", nil))
	}
}