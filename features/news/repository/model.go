package repository

import (
	"wakaf/features/news/domain"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255)"`
	Body    string `gorm:"type:text"`
	Picture string `gorm:"type:varchar(255)"`
}


func FromDomain(data domain.News) News {
	return News{
		Title: data.Title,
		Body: data.Body,
		Picture: data.Picture,
	}
}

func ToDomain(data News) domain.News {
	return domain.News{
		ID: data.ID,
		Title: data.Title,
		Body: data.Body,
		Picture: data.Picture,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}