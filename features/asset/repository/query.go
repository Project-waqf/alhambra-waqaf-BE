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

func (asset *AssetRepo) GetAll(status string, page int, sort string) ([]domain.Asset, int, int, int, error) {
	var res []Asset
	var countOnline, countDraft, countArchive int64

	if page != 0 {
		var offset int = 0
		offset = 8 * (page - 1)
		if status != "" {
			query := "SELECT * FROM assets WHERE status = ? AND deleted_at IS NULL ORDER BY created_at " + sort + " LIMIT ?, 8"
			if err := asset.db.Raw(query, status, offset).Find(&res).Error; err != nil {
				return []domain.Asset{}, 0, 0, 0, err
			}
		} else {
			order := "updated_at " +  sort
			if err := asset.db.Order(order).Limit(8).Offset(offset).Find(&res).Error; err != nil {
				return []domain.Asset{}, 0, 0, 0, err
			}
		}
	} else {
		order := "created_at " + sort
		if err := asset.db.Order(order).Find(&res).Error; err != nil {
			return []domain.Asset{}, 0, 0, 0, err
		}
	}

	if err := asset.db.Model(&Asset{}).Where("status = ?", "online").Count(&countOnline).Error; err != nil {
		return []domain.Asset{}, 0, 0, 0, err
	}

	if err := asset.db.Model(&Asset{}).Where("status = ?", "draft").Count(&countDraft).Error; err != nil {
		return []domain.Asset{}, 0, 0, 0, err
	}

	if err := asset.db.Model(&Asset{}).Where("status = ?", "archive").Count(&countArchive).Error; err != nil {
		return []domain.Asset{}, 0, 0, 0, err
	}

	return ToDomainGetAll(res), int(countOnline), int(countDraft), int(countArchive), nil
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
