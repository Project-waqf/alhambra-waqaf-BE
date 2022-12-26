package domain

import "time"

type Wakaf struct {
	ID        uint
	Title     string
	Category  string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCaseInterface interface {
	AddWakaf(input Wakaf) (Wakaf, error)
	GetAllWakaf() ([]Wakaf, error)
}

type RepoInterface interface {
	Insert(input Wakaf) (Wakaf, error)
	GetAllWakaf() ([]Wakaf, error)
}