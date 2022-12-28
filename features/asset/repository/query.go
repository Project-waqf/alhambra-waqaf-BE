package repository

import (
	"wakaf/features/asset/domain"

	"gorm.io/gorm"
)

type AssetRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepositoryInterface {
	return &AssetRepo{
		db: db,
	}
}

func (asset *AssetRepo) Insert(input domain.Asset) (domain.Asset, error) {
	data := FromDomainAdd(input)

	if err := asset.db.Create(&data).Last(&data).Error; err != nil {
		return domain.Asset{}, err
	}
	return ToDomainAdd(data), nil
}

func (asset *AssetRepo) GetAll() ([]domain.Asset, error) {
	var res []Asset

	if err := asset.db.Find(&res).Error; err != nil {
		return []domain.Asset{}, err
	}

	return ToDomainGetAll(res), nil
}

func (asset *AssetRepo) Get(id uint) (domain.Asset, error) {
	var res Asset

	if err := asset.db.First(&res, "id = ?", id).Error; err != nil {
		return domain.Asset{}, err
	}
	return ToDomainAdd(res), nil
}