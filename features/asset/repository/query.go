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

func (asset *AssetRepo) GetAll(status string) ([]domain.Asset, error) {
	var res []Asset

	if status == "online" {
		if err := asset.db.Where("status = 'online'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.Asset{}, err
		}
	} else if status == "draft" {
		if err := asset.db.Where("status = 'draft'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.Asset{}, err
		}
	} else if status == "archive" {
		if err := asset.db.Where("status = 'draft'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.Asset{}, err
		}
	} else {
		if err := asset.db.Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.Asset{}, err
		}
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

func (asset *AssetRepo) Edit(id uint, input domain.Asset) (domain.Asset, error) {
	data := FromDomainAdd(input)

	if err := asset.db.Where("id = ?", id).Updates(&data).Last(&data).Error; err != nil {
		return domain.Asset{}, err
	}
	data.ID = id
	return ToDomainAdd(data), nil
}

func (asset *AssetRepo) Delete(id uint) error {
	if err := asset.db.Delete(&Asset{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (asset *AssetRepo) ToOnline(id uint) error {

	if err := asset.db.First(&Asset{}, "id = ?", id).Error; err != nil {
		return err
	}

	if err := asset.db.Model(&Asset{}).Where("id = ?", id).Update("type", "online").Error; err != nil {
		return err
	}
	return nil
}

func (asset *AssetRepo) GetFileId(id uint) (string, error) {
	var res Asset

	if err := asset.db.Where("id = ?", id).First(&res).Error; err != nil {
		return "", err
	}
	return res.FileId, nil
}
