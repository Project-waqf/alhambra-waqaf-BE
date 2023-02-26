package delivery

import (
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

func ToDomainAdd(input WakafRequest) domain.Wakaf {
	var format = "2006-01-02"
	date, err := time.Parse(format + " 15:04:05", input.DueDate + " 23:59:59")
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
