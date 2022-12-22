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

	if err := news.db.Create(&cnv).Last(&cnv).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomainAddNews(cnv), nil
}

func (news *NewsRepository) GetAll() ([]domain.News, error) {
	var res []News

	if err := news.db.Find(&res).Error; err != nil {
		return []domain.News{}, err
	}

	return ToDomainGetAll(res), nil
}

func (news *NewsRepository) Get(id int) (domain.News, error) {
	var res News

	if err := news.db.Where("id = ?", id).First(&res).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomainGet(res), nil
}