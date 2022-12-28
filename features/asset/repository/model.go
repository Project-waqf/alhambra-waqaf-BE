package repository

import (
	"wakaf/features/asset/domain"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Name    string `gorm:"varchar(255)"`
	Picture string `gorm:"varchar(255)"`
	Detail  string `gorm:"varchar(255)"`
}

func FromDomainAdd(input domain.Asset) Asset {
	return Asset{
		Model: gorm.Model{ID: input.ID},
		Name: input.Name,
		Picture: input.Picture,
		Detail: input.Detail,
	}
}

func ToDomainAdd(input Asset) domain.Asset {
	return domain.Asset{
		ID: input.ID,
		Name: input.Name,
		Picture: input.Picture,
		Detail: input.Detail,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}