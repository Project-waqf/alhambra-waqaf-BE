package delivery

import (
	"wakaf/features/news/domain"
)

type NewsResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomainAddNews(data domain.News) NewsResponse {
	newCreated := data.CreatedAt.Format("Monday, 02-01-2006 T15:04:05")
	newUpdated := data.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
	return NewsResponse{
		ID: data.ID,
		Title: data.Title,
		Body: data.Body,
		Picture: data.Picture,
		CreatedAt: newCreated,
		UpdatedAt: newUpdated,
	}
}

func FromDomainGetAll(data []domain.News) []NewsResponse{
	var res []NewsResponse

	for _, v := range data {
		newCreated := v.CreatedAt.Format("Monday, 02-01-2006 T15:04:05")
		newUpdated := v.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
		res = append(res, NewsResponse{
			ID: v.ID,
			Title: v.Title,
			Body: v.Body,
			Picture: v.Picture,
			CreatedAt: newCreated,
			UpdatedAt: newUpdated,
		})
	}
	return res
}

func FromDOmainGet(data domain.News) NewsResponse {
	newCreated := data.CreatedAt.Format("Monday, 02-01-2006 T15:04:05")
	newUpdated := data.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05")
	return NewsResponse{
		ID: data.ID,
		Title: data.Title,
		Body: data.Body,
		Picture: data.Picture,
		CreatedAt: newCreated,
		UpdatedAt: newUpdated,
		}
}