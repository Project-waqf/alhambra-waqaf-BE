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
	cnv := FromDomain(input)

	if err := news.db.Create(&cnv).Last(&cnv).Error; err != nil {
		return domain.News{}, err
	}

	return ToDomain(cnv), nil
}