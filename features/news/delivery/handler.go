package delivery

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"wakaf/config"
	"wakaf/features/news/domain"
	"wakaf/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type NewsDelivery struct {
	NewsServices domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := NewsDelivery{
		NewsServices: data,
	}

	e.POST("/admin/news", handler.AddNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))               // ADD NEWS
	e.GET("/admin/news", handler.GetAllNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))             // GET ALL NEWS
	e.GET("/admin/news/:id_news", handler.GetSingleNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT))) // GET SINGLE NEWS
	e.PUT("/admin/news/:id_news", handler.UpdateNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.DELETE("/admin/news/:id_news", handler.DeleteNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
	e.PUT("admin/news/online/:id_news", handler.ToOnline(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))
}

func (news *NewsDelivery) AddNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input News

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		filename, err := helper.Upload(c, file, fileheader)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = filename
		cnv := ToDomainAddNews(input)
		res, err := news.NewsServices.AddNews(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		addResponse := FromDomainAddNews(res)
		return c.JSON(http.StatusOK, helper.Success("Add news Successfully", addResponse))
	}
}

func (news *NewsDelivery) GetAllNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := news.NewsServices.GetAll()
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}
		getAllResponse := FromDomainGetAll(res)
		return c.JSON(http.StatusOK, helper.Success("Get all news Successfully", getAllResponse))
	}
}

func (news *NewsDelivery) GetSingleNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		idTmp := c.Param("id_news")
		id, err := strconv.Atoi(idTmp)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		res, err := news.NewsServices.Get(id)
		if err != nil {
			log.Print(err.Error())
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, helper.Failed("Data not found"))
			}
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Get news successfully", FromDOmainGet(res)))
	}
}

func (news *NewsDelivery) UpdateNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input News
		err := c.Bind(&input)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		id := c.Param("id_news")
		cnvId, err:= strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		filename, err := helper.Upload(c, file, fileheader)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = filename
		cnvInput := ToDomainAddNews(input)
		res, err := news.NewsServices.UpdateNews(cnvId, cnvInput)
		if err != nil {	
			log.Print(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Update news successfully", FromDOmainGet(res)))
	}
}

func (news *NewsDelivery) DeleteNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_news")
		cnvId, err:= strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		_, err = news.NewsServices.Delete(cnvId)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Delete news successfully", nil))
	}
}

func (news *NewsDelivery) ToOnline() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_news")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		err = news.NewsServices.ToOnline(cnvId)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Set news to online successfully", nil))
	}
}
