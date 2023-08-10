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
	Status     string
	DueDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Donors struct {
	Name      string
	Fund      int `gorm:"column:gross_amount"`
	Doa       string
	CreatedAt time.Time
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

type Doa struct {
	Doa string
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
		Status:     input.Status,
		CreatedAt:  input.CreatedAt,
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
		Status:     input.Status,
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
			Status:     v.Status,
		})
	}
	return res
}

func ToDomainGet(input Wakaf, donors []Donors) domain.Wakaf {
	var newDonors []domain.Donors

	for _, v := range donors {
		var tmp = domain.Donors{
			Name:       v.Name,
			Fund:       v.Fund,
			Doa:        v.Doa,
			Created_at: input.CreatedAt,
		}
		newDonors = append(newDonors, tmp)
	}

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
		Donors:     newDonors,
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
