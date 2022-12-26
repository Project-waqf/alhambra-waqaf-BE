package repository

import (
	"wakaf/features/wakaf/domain"

	"gorm.io/gorm"
)

type WakafRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepoInterface {
	return &WakafRepo {
		db: db,
	}
}

func (wakaf *WakafRepo) Insert(input domain.Wakaf) (domain.Wakaf, error) {
	data := FromDomainAdd(input)

	if err := wakaf.db.Create(&data).Last(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}
	return ToDomainAdd(data), nil
}

func (wakaf *WakafRepo) GetAllWakaf() ([]domain.Wakaf, error) {
	var res []Wakaf

	if err := wakaf.db.Find(&res).Error; err != nil {
		return []domain.Wakaf{}, err
	}
	return ToDomainGetAll(res), nil
}