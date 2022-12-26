package repository

import (
	"wakaf/features/wakaf/domain"

	"gorm.io/gorm"
)

type Wakaf struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255)"`
	Category string `gorm:"type:varchar(255)"`
	Picture  string `gorm:"type:varchar(255)"`
}

func FromDomainAdd(input domain.Wakaf) Wakaf {
	return Wakaf{
		Model: gorm.Model{ID: input.ID},
		Title: input.Title,
		Category: input.Category,
		Picture: input.Picture,
	}
}

func ToDomainAdd(input Wakaf) domain.Wakaf {
	return domain.Wakaf{
		ID: input.ID,
		Title: input.Title,
		Category: input.Category,
		Picture: input.Picture,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ToDomainGetAll(input []Wakaf) []domain.Wakaf{
	var res []domain.Wakaf

	for _, v := range input {
		res = append(res, domain.Wakaf{
			ID: v.ID,
			Title: v.Title,
			Category: v.Category,
			Picture: v.Picture,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return res
}