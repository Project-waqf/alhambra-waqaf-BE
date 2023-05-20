package repository

import (
	"wakaf/features/partners/domain"

	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Name        string
	PictureName string
	Picture     string
	FileId      string
}

func FromDomainCreatePartner(input *domain.Partner) Partner {
	return Partner{
		Name:        input.Name,
		PictureName: input.PictureName,
		Picture:     input.Picture,
		FileId:      input.FileId,
	}
}

func ToDomainCreatePartner(input Partner) *domain.Partner {
	return &domain.Partner{
		Id:          input.ID,
		Name:        input.Name,
		PictureName: input.PictureName,
		Picture:     input.Picture,
		FileId:      input.FileId,
		CreatedAt:   input.CreatedAt,
		UpdateAt:    input.UpdatedAt,
	}
}

func ToDomainGetPartner(input Partner) *domain.Partner {
	return &domain.Partner{
		Id:          input.ID,
		Name:        input.Name,
		PictureName: input.PictureName,
		Picture:     input.Picture,
		FileId:      input.FileId,
		CreatedAt:   input.CreatedAt,
		UpdateAt:    input.UpdatedAt,
	}
}

func TodDomainGetListPartner(input []Partner) []*domain.Partner {
	var res []*domain.Partner

	for _, v := range input {
		tmp := ToDomainGetPartner(v)
		res = append(res, tmp)
	}
	return res
}
