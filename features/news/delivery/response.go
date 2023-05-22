package delivery

import (
	"wakaf/features/news/domain"
)

type NewsResponse struct {
	ID        uint   `json:"id_news"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Picture   string `json:"picture"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomainAddNews(data domain.News) NewsResponse {
	newCreated := data.CreatedAt.Format("Monday, 02-01-2006 T15:04:05")
	newUpdated := data.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
	return NewsResponse{
		ID:        data.ID,
		Title:     data.Title,
		Body:      data.Body,
		Picture:   data.Picture,
		Status:    data.Status,
		CreatedAt: newCreated,
		UpdatedAt: newUpdated,
	}
}

func FromDomainGetAll(data []domain.News) []NewsResponse {
	var res []NewsResponse

	for _, v := range data {
		newCreated := v.CreatedAt.Format("02 January 2006")
		newUpdated := v.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
		res = append(res, NewsResponse{
			ID:        v.ID,
			Title:     v.Title,
			Body:      v.Body,
			Picture:   v.Picture,
			Status:    v.Status,
			CreatedAt: newCreated,
			UpdatedAt: newUpdated,
		})
	}
	return res
}

func FromDomainGet(data domain.News) NewsResponse {
	newCreated := data.CreatedAt.Format("02 January 2006")
	newUpdated := data.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
	return NewsResponse{
		ID:        data.ID,
		Title:     data.Title,
		Body:      data.Body,
		Picture:   data.Picture,
		Status:    data.Status,
		CreatedAt: newCreated,
		UpdatedAt: newUpdated,
	}
}
