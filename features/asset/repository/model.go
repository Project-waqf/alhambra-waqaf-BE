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
	FileId  string
	Type    string
}

func FromDomainAdd(input domain.Asset) Asset {
	return Asset{
		Model:   gorm.Model{ID: input.ID},
		Name:    input.Name,
		Picture: input.Picture,
		Detail:  input.Detail,
		Type:    input.Type,
		FileId:  input.FileId,
	}
}

func ToDomainAdd(input Asset) domain.Asset {
	return domain.Asset{
		ID:        input.ID,
		Name:      input.Name,
		Picture:   input.Picture,
		Detail:    input.Detail,
		Type:      input.Type,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		FileId:    input.FileId,
	}
}

func ToDomainGetAll(input []Asset) []domain.Asset {
	var res []domain.Asset

	for _, v := range input {
		res = append(res, domain.Asset{
			ID:        v.ID,
			Name:      v.Name,
			Picture:   v.Picture,
			Detail:    v.Detail,
			Type:      v.Type,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return res
}
