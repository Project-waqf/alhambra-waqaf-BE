package delivery

import (
	"net/http"
	"strconv"
	"strings"
	"wakaf/config"
	"wakaf/features/news/domain"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type NewsDelivery struct {
	NewsServices domain.UseCaseInterface
}

var (
	logger = helper.Logger()
)

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := NewsDelivery{
		NewsServices: data,
	}

	e.POST("/admin/news", handler.AddNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))                 // ADD NEWS
	e.GET("/news", handler.GetAllNews())                                                                            // GET ALL NEWS
	e.GET("/news/:id_news", handler.GetSingleNews())                                                                // GET SINGLE NEWS
	e.PUT("/admin/news/:id_news", handler.UpdateNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))      // EDIT NEWS
	e.DELETE("/admin/news/:id_news", handler.DeleteNews(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))   // DELETE NEWS
	e.PUT("/admin/news/online/:id_news", handler.ToOnline(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT))) // FROM DRAFT TO ONLINE
}

func (news *NewsDelivery) AddNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input News

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			logger.Error("Error get file", zap.Error(err))
			fileId, filename, err := helper.Upload(c, file, fileheader, "news")
			if err != nil {
				logger.Error("Error upload image", zap.Error(err))
				filename = ""
				fileId = ""
			}
			input.Picture = filename
			input.FileId = fileId
		}

		res, err := news.NewsServices.AddNews(ToDomainAddNews(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		addResponse := FromDomainAddNews(res)
		return c.JSON(http.StatusOK, helper.Success("Add news Successfully", addResponse))
	}
}

func (news *NewsDelivery) GetAllNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		status := c.QueryParam("status")
		page := c.QueryParam("page")
		sort := c.QueryParam("sort")

		var pageCnv int
		if page != "" {
			cnv, err := strconv.Atoi(page)
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.Failed("Error query parameter"))
			}
			pageCnv = cnv
		}
		res, countOnline, countDraft, countArchive, err := news.NewsServices.GetAll(status, pageCnv, sort)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Something error in server"))
		}
		getAllResponse := FromDomainGetAll(res)
		return c.JSON(http.StatusOK, helper.SuccessGetAll("Get all news Successfully", getAllResponse, countOnline, countDraft, countArchive))
	}
}

func (news *NewsDelivery) GetSingleNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		idTmp := c.Param("id_news")
		id, err := strconv.Atoi(idTmp)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		res, err := news.NewsServices.Get(id)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, helper.Failed("Data not found"))
			}
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Get news successfully", FromDomainGet(res)))
	}
}

func (news *NewsDelivery) UpdateNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input News

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		id := c.Param("id_news")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileIdDb, err := news.NewsServices.GetFileId(cnvId)
			if err != nil {
				logger.Error("Failed to get fileId", zap.Error(err))
				return c.JSON(http.StatusNotFound, helper.Failed("Failed to get fileId"))
			} else if err == nil && fileIdDb == "" {
				fileId, filename, err := helper.Upload(c, file, fileheader, "news")
				if err != nil {
					return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
				}
				input.FileId = fileId
				input.Picture = filename
			} else {
				err = helper.Delete(fileIdDb)
				if err != nil {
					logger.Error("Failed delete image in imagekit", zap.Error(err))
					return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
				}
				fileId, filename, err := helper.Upload(c, file, fileheader, "news")
				if err != nil {
					return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
				}
				input.FileId = fileId
				input.Picture = filename
			}
		}

		res, err := news.NewsServices.UpdateNews(cnvId, ToDomainAddNews(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Update news successfully", FromDomainGet(res)))
	}
}

func (news *NewsDelivery) DeleteNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_news")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		fileIdDb, err := news.NewsServices.GetFileId(cnvId)
		if err == nil && fileIdDb != "" {
			err = helper.Delete(fileIdDb)
			if err != nil {
				logger.Error("Error delete file in imagekit", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
			}
		}

		_, err = news.NewsServices.Delete(cnvId)
		if err != nil {
			logger.Error("Error in usecase", zap.Error(err))
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, helper.Failed("Data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
		}
		return c.JSON(http.StatusOK, helper.Success("Delete news successfully", nil))
	}
}

func (news *NewsDelivery) ToOnline() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_news")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		err = news.NewsServices.ToOnline(cnvId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusCreated, helper.Success("Set news to online successfully", nil))
	}
}
