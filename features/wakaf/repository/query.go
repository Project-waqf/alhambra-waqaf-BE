package repository

import (
	"database/sql"
	"fmt"
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

func (wakaf *WakafRepo) GetAllWakaf(category string, page int, isUser bool, sort, filter, status string) ([]domain.Wakaf, int, int, int, error) {
	var res []Wakaf
	var resWithStatus []Wakaf
	var countOnline, countDraft, countArchive int64

	today := time.Now()

	var offset int = 0
	if page != 1 {
		offset = 9 * (page - 1)
	}

	if category != "" {
		if page != 0 {
			if isUser {
				if err := wakaf.db.Raw("SELECT * FROM wakafs WHERE due_date >= NOW() AND collected != fund_target AND status = ? AND category = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT ?, 9", status, category, offset).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			} else {
				if err := wakaf.db.Where("category = ?", category).Order("updated_at DESC").Limit(9).Offset(offset).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			}
		} else {
			if isUser {
				if err := wakaf.db.Where("category = ? AND due_date >= ?", category, today).Order("updated_at desc").Limit(9).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			} else {
				if err := wakaf.db.Where("category = ?", category, today).Order("updated_at desc").Limit(9).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			}
		}
	} else if filter != "" {
		if filter == "aktif" {
			query := "SELECT * FROM wakafs WHERE due_date >= NOW() AND collected < fund_target AND status = 'online' AND deleted_at IS NULL ORDER BY created_at " + sort
			if err := wakaf.db.Raw(query).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, 0, 0, err
			}
		} else if filter == "complete" {
			query := "SELECT * FROM wakafs WHERE collected > 0 AND collected = fund_target AND status = 'online' AND deleted_at IS NULL ORDER BY created_at " + sort
			if err := wakaf.db.Raw(query).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, 0, 0, err
			}
		} else {
			query := "SELECT * FROM wakafs WHERE collected < fund_target AND due_date < NOW() AND status = 'online' AND deleted_at IS NULL ORDER BY created_at " + sort
			if err := wakaf.db.Raw(query).Find(&res).Error; err != nil {
				return []domain.Wakaf{}, 0, 0, 0, err
			}
		}
	} else {
		if page != 0 {
			if isUser {
				if err := wakaf.db.Raw("SELECT * FROM wakafs WHERE due_date >= NOW() AND deleted_at IS NULL ORDER BY RAND() DESC LIMIT ?, 9", offset).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			} else {
				query := "SELECT * FROM wakafs WHERE status = ? AND deleted_at IS NULL ORDER BY created_at " + sort + "  LIMIT ?, 9"
				if err := wakaf.db.Raw(query, status, offset).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			}
		} else {
			if isUser {
				if err := wakaf.db.Raw("SELECT * FROM wakafs WHERE due_date >= ? AND status = ? AND deleted_at IS NULL ORDER BY created_at DESC", today, status).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			} else {
				query := "SELECT * FROM wakafs WHERE deleted_at IS NULL ORDER BY created_at " + sort
				if err := wakaf.db.Raw(query).Find(&res).Error; err != nil {
					return []domain.Wakaf{}, 0, 0, 0, err
				}
			}
		}
	}

	if category != "" && isUser {
		if err := wakaf.db.Model(&Wakaf{}).Where("status = ? AND category = ? AND due_date >= ? AND fund_target != collected", status, category, today).Count(&countOnline).Error; err != nil {
			return []domain.Wakaf{}, 0, 0, 0, err
		}
	} else {
		if err := wakaf.db.Model(&Wakaf{}).Where("status = ?", "online").Count(&countOnline).Error; err != nil {
			return []domain.Wakaf{}, 0, 0, 0, err
		}
	}

	if err := wakaf.db.Model(&Wakaf{}).Where("status = ?", "draft").Count(&countDraft).Error; err != nil {
		return []domain.Wakaf{}, 0, 0, 0, err
	}

	if err := wakaf.db.Model(&Wakaf{}).Where("status = ?", "archive").Count(&countArchive).Error; err != nil {
		return []domain.Wakaf{}, 0, 0, 0, err
	}

	if status != "" {
		for i := 0; i < len(res); i++ {
			if isUser {
				if res[i].Status == status && res[i].Collected != res[i].FundTarget {
					resWithStatus = append(resWithStatus, res[i])
				}
			} else {
				if res[i].Status == status {
					resWithStatus = append(resWithStatus, res[i])
				}
			}
		}

		return ToDomainGetAll(resWithStatus), int(countOnline), int(countDraft), int(countArchive), nil
	}

	return ToDomainGetAll(res), int(countOnline), int(countDraft), int(countArchive), nil
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
	return ToDomainGet(data, nil), nil
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
	var donors []Donors

	if err := wakaf.db.Where("id = ?", id).First(&data).Error; err != nil {
		return domain.Wakaf{}, err
	}

	if err := wakaf.db.Table("donors").Select("name, gross_amount, doa").Where("id_wakaf = ?", id).Limit(10).Scan(&donors).Error; err != nil {
		return domain.Wakaf{}, err
	}

	return ToDomainGet(data, donors), nil
}

func (Wakaf *WakafRepo) PayWakaf(input domain.PayWakaf) (domain.PayWakaf, error) {
	data := FromDomainPaywakaf(input)
	var res Donor
	if err := Wakaf.db.Model(&Donor{}).Create(&data).Last(&res).Error; err != nil {
		return domain.PayWakaf{}, err
	}

	return ToDomainPayment(res), nil
}

func (wk *WakafRepo) UpdatePayment(input domain.PayWakaf) (domain.PayWakaf, error) {
	var res Donor
	var id_wakaf int

	if err := wk.db.Model(&Donor{}).Select("id_wakaf").Where("order_id", input.OrderId).Scan(&id_wakaf).Error; err != nil {
		return domain.PayWakaf{}, err
	}

	fmt.Println("[DEBUG] id_wakaf : ", id_wakaf)

	if err := wk.db.Exec("UPDATE wakafs SET collected = collected + @gross_amount WHERE id = @id_wakaf", sql.Named("gross_amount", input.GrossAmount), sql.Named("id_wakaf", id_wakaf)).Error; err != nil {
		return domain.PayWakaf{}, err
	}

	if err := wk.db.Table("donors").Where("order_id = ?", input.OrderId).Updates(Donor{Status: input.Status, PaymentType: input.PaymentType}).Last(&res).Error; err != nil {
		return domain.PayWakaf{}, err
	}
	return ToDomainPayment(res), nil
}

func (wk *WakafRepo) Search(input string) ([]domain.Wakaf, int, int, int, error) {
	var res []Wakaf
	var countOnline, countDraft, countArchive int

	if input != "" {
		if err := wk.db.Where("title like ?", "%"+input+"% AND due_date >= NOW() AND collected < fund_target").Find(&res).Error; err != nil {
			return []domain.Wakaf{}, 0, 0, 0, err
		}
	} else {
		if err := wk.db.Exec("SELECT * FROM wakafs WHERE status = online AND collected < fund_target AND due_date >= NOW() AND deleted_at IS NULL ORDER BY created_at desc").Find(&res).Error; err != nil {
			return []domain.Wakaf{}, 0, 0, 0, err
		}
	}

	for _, v := range res {
		if v.Status == "online" {
			countOnline += 1
		} else if v.Status == "draft" {
			countDraft += 1
		} else {
			countArchive += 1
		}
	}

	return ToDomainGetAll(res), countOnline, countDraft, countArchive, nil
}

func (wk *WakafRepo) GetSummary() (int, int, int, error) {
	var count, sum, wakif int64

	if err := wk.db.Model(&Wakaf{}).Count(&count).Error; err != nil {
		return 0, 0, 0, err
	}

	if err := wk.db.Raw("SELECT sum(collected) as sum_collected FROM wakafs").Scan(&sum).Error; err != nil {
		return 0, 0, 0, err
	}

	if err := wk.db.Model(&Donor{}).Where("status = ?", "settlement").Count(&wakif).Error; err != nil {
		return 0, 0, 0, err
	}
	return int(count), int(sum), int(wakif), nil
}
