package domain

import "time"

type News struct {
	ID        uint
	Title     string
	Body      string
	Picture   string
	Type      string
	FileId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}

type UseCaseInterface interface {
	AddNews(input News) (News, error)
	GetAll(status string, page int) ([]News, int, error)
	Get(id int) (News, error)
	UpdateNews(id int, input News) (News, error)
	Delete(id int) (News, error)
	ToOnline(id int) error
	GetFileId(id int) (string, error)
}

type RepoInterface interface {
	Insert(input News) (News, error)
	GetAll(status string, page int) ([]News, int, error)
	Get(id int) (News, error)
	Edit(id int, input News) (News, error)
	Delete(id int) (News, error)
	ToOnline(id int) error
	GetFileId(id int) (string, error)
}
