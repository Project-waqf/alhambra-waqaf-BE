package repository

import (
	"wakaf/features/partners/domain"

	"gorm.io/gorm"
)

type PartnerRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepoInterface {
	return &PartnerRepo{
		db: db,
	}
}
