package delivery

import "wakaf/features/news/domain"

type News struct {
	Title   string `json:"title" form:"title"`
	Body    string `json:"body" form:"body"`
	Picture string `json:"picture" form:"picture"`
	FileId  string
	Status string `json:"status" form:"status"`
}

func ToDomainAddNews(data News) domain.News {
	return domain.News{
		Title:   data.Title,
		Body:    data.Body,
		Picture: data.Picture,
		FileId: data.FileId,
		Status: data.Status,
	}
}
