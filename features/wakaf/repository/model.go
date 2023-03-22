package repository

import (
	"time"
	"wakaf/features/wakaf/domain"

	"gorm.io/gorm"
)

type Wakaf struct {
	gorm.Model
	Title      string `gorm:"type:varchar(255)"`
	Category   string `gorm:"type:varchar(255)"`
	Picture    string `gorm:"type:varchar(255)"`
	Detail     string
	Collected  int
	FundTarget int
	FileId     string
	DueDate    *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Donor struct {
	gorm.Model
	IdWakaf     uint
	Name        string
	Doa         string
	GrossAmount int
	Status      string
	PaymentType string
	OrderId     string
}

func FromDomainAdd(input domain.Wakaf) Wakaf {
	return Wakaf{
		Model:      gorm.Model{ID: input.ID},
		Title:      input.Title,
		Category:   input.Category,
		Picture:    input.Picture,
		FileId:     input.FileId,
		Detail:     input.Detail,
		Collected:  input.Collected,
		FundTarget: input.FundTarget,
		DueDate:    input.DueDate,
	}
}

func FromDomainPaywakaf(input domain.PayWakaf) Donor {
	return Donor{
		IdWakaf:     uint(input.IdWakaf),
		Name:        input.Name,
		GrossAmount: input.GrossAmount,
		Doa:         input.Doa,
		PaymentType: input.PaymentType,
		OrderId:     input.OrderId,
		Status:      input.Status,
	}
}

func ToDomainAdd(input Wakaf) domain.Wakaf {
	return domain.Wakaf{
		ID:         input.ID,
		Title:      input.Title,
		Category:   input.Category,
		Picture:    input.Picture,
		CreatedAt:  input.CreatedAt,
		UpdatedAt:  input.UpdatedAt,
		FileId:     input.FileId,
		Detail:     input.Detail,
		FundTarget: input.FundTarget,
		DueDate:    input.DueDate,
		Collected:  input.Collected,
	}
}

func ToDomainGetAll(input []Wakaf) []domain.Wakaf {
	var res []domain.Wakaf

	for _, v := range input {
		res = append(res, domain.Wakaf{
			ID:         v.ID,
			Title:      v.Title,
			Category:   v.Category,
			Picture:    v.Picture,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			Detail:     v.Detail,
			FundTarget: v.FundTarget,
			DueDate:    v.DueDate,
			Collected:  v.Collected,
		})
	}
	return res
}

func ToDomainGet(input Wakaf) domain.Wakaf {
	return domain.Wakaf{
		ID:         input.ID,
		Title:      input.Title,
		Category:   input.Category,
		Picture:    input.Picture,
		CreatedAt:  input.UpdatedAt,
		UpdatedAt:  input.UpdatedAt,
		Detail:     input.Detail,
		FundTarget: input.FundTarget,
		DueDate:    input.DueDate,
		Collected:  input.Collected,
	}
}

func ToDomainPayment(input Donor) domain.PayWakaf {
	return domain.PayWakaf{
		IdWakaf:     int(input.IdWakaf),
		Name:        input.Name,
		GrossAmount: input.GrossAmount,
		Doa:         input.Doa,
		OrderId:     input.OrderId,
		Status:      input.Status,
		PaymentType: input.PaymentType,
		CreatedAt:   input.CreatedAt,
	}
}
