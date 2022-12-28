package domain

import "time"

type Asset struct {
	ID        uint
	Name      string
	Picture   string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsecaseInterface interface {
	AddAsset(input Asset) (Asset, error)
	GetAllAsset() ([]Asset, error)
}

type RepositoryInterface interface {
	Insert(input Asset) (Asset, error)
	GetAll() ([]Asset, error)
}
