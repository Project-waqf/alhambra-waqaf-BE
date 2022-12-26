package repository

import (
	"log"
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

func (news *NewsRepository) Edit(id int, input domain.News) (domain.News, error) {
	data := FromDomainAddNews(input)

	if err := news.db.Model(&News{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		log.Print(err.Error())
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