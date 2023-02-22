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
	Status  string 
	FileId  string
}

func FromDomainAddNews(data domain.News) News {
	return News{
		Title:   data.Title,
		Body:    data.Body,
		Picture: data.Picture,
		FileId:  data.FileId,
		Status:  data.Status,
	}
}

func ToDomainAddNews(data News) domain.News {
	return domain.News{
		ID:        data.ID,
		Title:     data.Title,
		Body:      data.Body,
		Picture:   data.Picture,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		FileId:    data.FileId,
		Status:    data.Status,
	}
}

func ToDomainGetAll(data []News) []domain.News {
	var res []domain.News

	for _, v := range data {
		res = append(res, domain.News{
			ID:        v.ID,
			Title:     v.Title,
			Body:      v.Body,
			Picture:   v.Picture,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return res
}

func ToDomainGet(data News) domain.News {
	return domain.News{
		ID:        data.ID,
		Title:     data.Title,
		Body:      data.Body,
		Picture:   data.Picture,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		FileId: data.FileId,
	}
}
