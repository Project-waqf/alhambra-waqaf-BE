package domain

import "time"

type News struct {
	ID        uint
	Title     string
	Body      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCaseInterface interface {
	AddNews(input News) (News, error)
	GetAll() ([]News, error)
	Get(id int) (News, error)
	UpdateNews(id int, input News) (News, error)
	Delete(id int) (News,error)
}

type RepoInterface interface {
	Insert(input News) (News, error)
	GetAll() ([]News, error)
	Get(id int) (News, error)
	Edit(id int, input News) (News, error)
	Delete(id int) (News, error)
}