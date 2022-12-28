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

func (wakaf *WakafRepo) Edit(id uint, input domain.Wakaf) (domain.Wakaf, error) {
	data := FromDomainAdd(input)

	if err := wakaf.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}

	data.ID = id
	return ToDomainAdd(data), nil
}

func (wakaf *WakafRepo) Delete(id uint) (domain.Wakaf, error) {
	var data Wakaf

	if err := wakaf.db.Where("id = ?", id).First(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}

	if err := wakaf.db.Delete(&Wakaf{}, "id = ?", id).Error; err != nil {
		return domain.Wakaf{}, err
	}
	return ToDomainGet(data), nil
}