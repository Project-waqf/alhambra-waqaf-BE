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