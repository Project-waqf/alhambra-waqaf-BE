package repository

import (
	"time"
	"wakaf/features/wakaf/domain"

	"gorm.io/gorm"
)

type WakafRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepoInterface {
	return &WakafRepo{
		db: db,
	}
}

func (wakaf *WakafRepo) Insert(input domain.Wakaf) (domain.Wakaf, error) {
	input.Collected = 0
	data := FromDomainAdd(input)

	if err := wakaf.db.Create(&data).Last(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}
	return ToDomainAdd(data), nil
}

func (wakaf *WakafRepo) GetAllWakaf(category string, page int) ([]domain.Wakaf, int, error) {
	var res []Wakaf
	var count int64

	today := time.Now()

	var offset int = 0
	if page != 1 {
		offset = 9 * (page - 1)
	}

	if category != "" {
		if page != 0 {
			if err := wakaf.db.Where("category = ? AND due_date >= 0", category).Order("created_at DESC").Limit(9).Offset(offset).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, err
			}
		} else {
			if err := wakaf.db.Where("category = ? AND due_date >= ?", category, today).Order("created_at desc").Limit(9).Offset(offset).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, err
			}
		}
	} else {
		if page != 0 {
			if err := wakaf.db.Where("due_date >= ?", today).Order("created_at desc").Limit(9).Offset(offset).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, err
			}
		} else {
			if err := wakaf.db.Where("due_date >= ?", today).Order("created_at desc").Limit(9).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, err
			}
		}
	}

	if err := wakaf.db.Model(&Wakaf{}).Count(&count).Error; err != nil {
		return []domain.Wakaf{}, 0, err
	}
	return ToDomainGetAll(res), int(count), nil
}

func (wakaf *WakafRepo) Edit(id uint, input domain.Wakaf) (domain.Wakaf, error) {
	data := FromDomainAdd(input)

	if err := wakaf.db.Where("id = ?", id).Updates(&data).Last(&data).Error; err != nil {
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

func (wakaf *WakafRepo) GetFileId(id uint) (string, error) {
	var data Wakaf

	if err := wakaf.db.Where("id = ?", id).First(&data).Error; err != nil {
		return "", err
	}
	return data.FileId, nil
}

func (wakaf *WakafRepo) GetSingleWakaf(id uint) (domain.Wakaf, error) {
	var data Wakaf

	if err := wakaf.db.Where("id = ?", id).First(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}

	return ToDomainGet(data), nil
}

func (Wakaf *WakafRepo) PayWakaf(input domain.PayWakaf) (domain.PayWakaf, error) {
	data := FromDomainPaywakaf(input)
	var res Donor

	if err := Wakaf.db.Model(&Donor{}).Create(&data).Last(&res).Error; err != nil {
		return domain.PayWakaf{}, err
	}
	return ToDomainPayment(res), nil
}

func (Wakaf *WakafRepo) UpdatePayment(input domain.PayWakaf) (domain.PayWakaf, error) {
	data := FromDomainPaywakaf(input)
	var res Donor

	if err := Wakaf.db.Raw("UPDATE wakafs set collected = collected + ? WHERE id = ?", input.GrossAmount, input.IdWakaf).Error; err != nil {
		return domain.PayWakaf{}, nil
	}

	if err := Wakaf.db.Model(&Donor{}).Where("order_id = ?", input.OrderId).Update("status", data.Status).Update("payment_type", input.PaymentType).Last(&res).Error; err != nil {
		return domain.PayWakaf{}, err
	}
	return ToDomainPayment(res), nil
}
