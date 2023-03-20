package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wakaf/config"
	"wakaf/features/wakaf/domain"
	"wakaf/pkg/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type WakafDelivery struct {
	WakafService domain.UseCaseInterface
}

func New(e *echo.Echo, data domain.UseCaseInterface) {
	handler := WakafDelivery{
		WakafService: data,
	}

	e.POST("/admin/wakaf", handler.AddWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))                // INSERT WAKAF
	e.GET("/wakaf", handler.GetAllWakaf())                                                                           // GET ALL WAKAF
	e.PUT("/admin/wakaf/:id_wakaf", handler.UpdateWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT)))    // UPDATE WAKAF
	e.DELETE("/admin/wakaf/:id_wakaf", handler.DeleteWakaf(), middleware.JWT([]byte(config.Getconfig().SECRET_JWT))) // DELETE WAKAF
	e.GET("/wakaf/:id_wakaf", handler.GetSingleWakaf())
	e.POST("/wakaf/pay", handler.PayWakaf())
	e.POST("/wakaf/payment/callback", handler.PaymentCallback())
}

var (
	logger = helper.Logger()
)

func (wakaf *WakafDelivery) AddWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WakafRequest

		err := c.Bind(&input)
		if err != nil {
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err != nil {
			logger.Error("Error get picture", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		fileId, dest, err := helper.Upload(c, file, fileheader, "wakaf")
		if err != nil {
			logger.Error("Error upload images", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		input.Picture = dest
		input.FileId = fileId
		res, err := wakaf.WakafService.AddWakaf(ToDomainAdd(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Add wakaf successfully", FromDomainAdd(res)))
	}
}

func (wakaf *WakafDelivery) GetAllWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		category := c.QueryParam("category")
		page := c.QueryParam("page")
		cnvPage, err := strconv.Atoi(page)
		if err != nil {
			logger.Error("Failed to convert query param page")
		}

		res, count, err := wakaf.WakafService.GetAllWakaf(category, cnvPage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.SuccessGetAll("Get all wakaf successfully", FromDomainGetAll(res), count))
	}
}

func (wakaf *WakafDelivery) UpdateWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input WakafRequest

		id := c.Param("id_wakaf")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			logger.Error("Failed convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		err = c.Bind(&input)
		if err != nil {
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		file, fileheader, err := c.Request().FormFile("picture")
		if err == nil {
			fileIdDb, err := wakaf.WakafService.GetFileId(uint(cnvId))
			if err != nil {
				logger.Error("Failed to get fileId", zap.Error(err))
				return c.JSON(http.StatusNotFound, helper.Failed("Failed to get fileId"))
			}
			err = helper.Delete(fileIdDb)
			if err != nil {
				logger.Error("Failed delete image in imagekit", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
			}

			fileId, fileName, err := helper.Upload(c, file, fileheader, "wakaf")
			if err != nil {
				logger.Error("Error upload image", zap.Error(err))
				return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
			}
			input.FileId = fileId
			input.Picture = fileName
		}

		res, err := wakaf.WakafService.UpdateWakaf(uint(cnvId), ToDomainAdd(input))
		if err != nil {
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
			logger.Error("Failed convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		fileIdDb, err := wakaf.WakafService.GetFileId(uint(cnvId))
		if err != nil {
			logger.Error("Failed to get fileId", zap.Error(err))
			return c.JSON(http.StatusNotFound, helper.Failed("Failed to get fileId"))
		}
		err = helper.Delete(fileIdDb)
		if err != nil {
			logger.Error("Failed delete image in imagekit", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, helper.Failed("Failed to update"))
		}

		_, err = wakaf.WakafService.DeleteWakaf(uint(cnvId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, helper.Success("Delete wakaf successfully", nil))
	}
}

func (wakaf *WakafDelivery) GetSingleWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id_wakaf")
		cnvId, err := strconv.Atoi(id)
		if err != nil {
			logger.Error("Error when convert id", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		res, err := wakaf.WakafService.GetSingleWakaf(uint(cnvId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}

		return c.JSON(http.StatusOK, FromDomainGet(res))
	}
}

func (wakaf *WakafDelivery) PayWakaf() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input PayWakafReq

		if err := c.Bind(&input); err != nil {
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}

		res, err := wakaf.WakafService.PayWakaf(ToDomainPayWakaf(input))
		if err != nil {
			if strings.Contains(err.Error(), "not found") || err.Error() == "completed" {
				return c.JSON(http.StatusNotFound, helper.Failed("Funding has completed"))
			}
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}
		return c.JSON(http.StatusOK, FromDomainPaywakaf(res))
	}
}

func (wakaf *WakafDelivery) PaymentCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CallbackMidtrans
		
		if err := c.Bind(&input); err != nil {
			logger.Error("Error bind data", zap.Error(err))
			return c.JSON(http.StatusBadRequest, helper.Failed("Error input"))
		}
		
		fmt.Println("[DEBUG] Data Callback", input)

		// Fraud Check 
		if input.FraudStatus == "deny" || input.FraudStatus == "challenge" {
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))

			err := wakaf.WakafService.DenyTransaction(input.OrderId)
			if err != nil {
				logger.Error("Failed to deny transaction")
			}
			logger.Error("Failed payment")
			return c.JSON(http.StatusOK, helper.Failed("Failed transaction"))
		}

		// Transaction Status Check
		switch input.TransactionStatus {
		case "pending":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		case "deny":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		case "cancel":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		case "expire":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		case "refund":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		case "authorize":
			logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
			return c.JSON(http.StatusOK, helper.Failed("Payment"+input.TransactionStatus))
		}


		res, err := wakaf.WakafService.UpdatePayment(ToDomainCallback(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Failed("Something error in server"))
		}

		logger.Info("Payment "+input.TransactionStatus, zap.Any("Order Id", input.OrderId))
		return c.JSON(http.StatusOK, helper.Success("Update payment successfull", res))
	}
}
