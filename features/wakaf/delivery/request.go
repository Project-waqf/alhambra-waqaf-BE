package delivery

import (
	"strconv"
	"strings"
	"time"
	"wakaf/features/wakaf/domain"

	"go.uber.org/zap"
)

type WakafRequest struct {
	ID         uint
	Title      string `json:"title" form:"title"`
	Category   string `json:"category" form:"category"`
	Detail     string `json:"detail" form:"detail"`
	Picture    string `json:"picture" form:"picture"`
	FundTarget int    `json:"fund_target" form:"fund_target"`
	DueDate    string `json:"due_date" form:"due_date"`
	FileId     string
}

type PayWakafReq struct {
	IdWakaf     int    `json:"id_wakaf" form:"id_wakaf"`
	Name        string `json:"name" form:"name"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Doa         string `json:"doa" form:"doa"`
}

type CallbackMidtrans struct {
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
}

func ToDomainAdd(input WakafRequest) domain.Wakaf {
	var format = "2006-01-02"
	date, err := time.Parse(format+" 15:04:05", input.DueDate+" 23:59:59")
	if err != nil {
		logger.Error("Error parse due date", zap.Error(err))
		return domain.Wakaf{}
	}

	return domain.Wakaf{
		ID:         input.ID,
		Title:      input.Title,
		Category:   input.Category,
		Picture:    input.Picture,
		FileId:     input.FileId,
		Detail:     input.Detail,
		FundTarget: input.FundTarget,
		DueDate:    &date,
	}
}

func ToDomainPayWakaf(input PayWakafReq) domain.PayWakaf {
	return domain.PayWakaf{
		IdWakaf:     input.IdWakaf,
		Name:        input.Name,
		GrossAmount: input.GrossAmount,
		Doa:         input.Doa,
	}
}

func ToDomainCallback(input CallbackMidtrans) domain.PayWakaf {
	split := strings.Split(input.GrossAmount, ".")
	cnv, _ := strconv.Atoi(split[0])
	cnvNew := cnv / 2
	return domain.PayWakaf{
		OrderId:     input.OrderId,
		PaymentType: input.PaymentType,
		Status:      input.TransactionStatus,
		GrossAmount: cnvNew,
	}
}
