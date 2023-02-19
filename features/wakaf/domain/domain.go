package domain

import "time"

type Wakaf struct {
	ID        uint
	Title     string
	Category  string
	Picture   string
	FileId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCaseInterface interface {
	AddWakaf(input Wakaf) (Wakaf, error)
	GetAllWakaf() ([]Wakaf, error)
	UpdateWakaf(id uint, input Wakaf) (Wakaf, error)
	DeleteWakaf(id uint) (Wakaf, error)
}

type RepoInterface interface {
	Insert(input Wakaf) (Wakaf, error)
	GetAllWakaf() ([]Wakaf, error)
	Edit(id uint, input Wakaf) (Wakaf, error)
	Delete(id uint) (Wakaf, error)
}
