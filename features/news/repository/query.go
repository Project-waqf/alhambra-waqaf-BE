package repository

import (
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

func (news *NewsRepository) GetAll(status string) ([]domain.News, error) {
	var res []News

	if status == "online" {
		if err := news.db.Where("status = 'online'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.News{}, err
		}
	} else if status == "draft" {
		if err := news.db.Where("status = 'draft'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.News{}, err
		}
	} else if status == "archive" {
		if err := news.db.Where("status = 'archive'").Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.News{}, err
		}
	} else {
		if err := news.db.Order("updated_at DESC").Find(&res).Error; err != nil {
			return []domain.News{}, err
		}
	}

	return ToDomainGetAll(res), nil
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