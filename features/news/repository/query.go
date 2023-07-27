package repository

import (
	"strings"
	"wakaf/features/news/domain"

	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepoInterface {
	return &NewsRepository{
		db: db,
	}
}

func (news *NewsRepository) Insert(input domain.News) (domain.News, error) {
	cnv := FromDomainAddNews(input)

	if err := news.db.Create(&cnv).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomainAddNews(cnv), nil
}

func (news *NewsRepository) GetAll(status string, page int, sort string) ([]domain.News, int, int, int, error) {
	var res []News
	var countOnline, countDraft, countArchive int64

	if page != 0 {
		var offset int = 0
		offset = 10 * (page - 1)
		if status == "online" {
			order := "created_at " + sort
			if err := news.db.Where("status = 'online'").Order(order).Limit(10).Offset(offset).Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else if status == "draft" {
			order := "created_at " + sort
			if err := news.db.Where("status = 'draft'").Order(order).Limit(10).Offset(offset).Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else if status == "archive" {
			order := "created_at " + sort
			if err := news.db.Where("status = 'archive'").Order(order).Limit(10).Offset(offset).Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else {
			order := "created_at " + sort
			if err := news.db.Order(order).Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		}
	} else {
		if status == "online" {
			if err := news.db.Where("status = 'online'").Order("updated_at DESC").Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else if status == "draft" {
			if err := news.db.Where("status = 'draft'").Order("updated_at DESC").Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else if status == "archive" {
			if err := news.db.Where("status = 'archive'").Order("updated_at DESC").Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		} else {
			if err := news.db.Order("updated_at DESC").Find(&res).Error; err != nil {
				return []domain.News{}, 0, 0, 0, err
			}
		}
	}

	if err := news.db.Model(&News{}).Where("status = ?", "online").Count(&countOnline).Error; err != nil {
		return []domain.News{}, 0, 0, 0, err
	}

	if err := news.db.Model(&News{}).Where("status = ?", "draft").Count(&countDraft).Error; err != nil {
		return []domain.News{}, 0, 0, 0, err
	}

	if err := news.db.Model(&News{}).Where("status = ?", "archive").Count(&countArchive).Error; err != nil {
		return []domain.News{}, 0, 0, 0, err
	}

	return ToDomainGetAll(res), int(countOnline), int(countDraft), int(countArchive), nil
}

func (news *NewsRepository) Get(id int) (domain.News, error) {
	var res News

	if err := news.db.Where("id = ? and status = 'online'", id).First(&res).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomainGet(res), nil
}

func (news *NewsRepository) Edit(id int, input domain.News) (domain.News, error) {
	data := FromDomainAddNews(input)

	if err := news.db.Model(&News{}).Where("id = ?", id).Updates(&data).Last(&data).Error; err != nil {
		return domain.News{}, err
	}

	data.ID = uint(id)
	return ToDomainAddNews(data), nil
}

func (news *NewsRepository) Delete(id int) (domain.News, error) {
	var res News

	if err := news.db.Where("id = ?", id).First(&res).Error; err != nil {
		return domain.News{}, err
	}

	if err := news.db.Model(&News{}).Delete("id = ?", id).Error; err != nil {
		return domain.News{}, err
	}
	return ToDomainGet(res), nil
}

func (news *NewsRepository) ToOnline(id int) error {

	if err := news.db.Model(&News{}).Where("id = ?", id).Error; err != nil {
		return err
	}

	if err := news.db.Model(&News{}).Where("id = ?", id).Update("type", "online").Error; err != nil {
		return err
	}
	return nil
}

func (news *NewsRepository) GetFileId(id int) (string, error) {
	var res News

	if err := news.db.Where("id = ?", id).First(&res).Error; err != nil {
		return "", err
	}

	return ToDomainGet(res).FileId, nil
}

func (news *NewsRepository) GetBySlug(slug string) (domain.News, error) {
	var res News

	newSlug := strings.ReplaceAll(slug, "-", " ")
	if strings.Contains(slug, "_") {
		newSlug = strings.ReplaceAll(newSlug, "_", "-")
	}

	if err := news.db.Where("title = ? and status = 'online'", newSlug).First(&res).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomainGet(res), nil
}
