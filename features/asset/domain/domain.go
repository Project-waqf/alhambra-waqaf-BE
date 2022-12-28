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
	GetAsset(id uint) (Asset, error)
	UpdateAsset(id uint, input Asset) (Asset, error)
	DeleteAsset(id uint) (error)
}

type RepositoryInterface interface {
	Insert(input Asset) (Asset, error)
	GetAll() ([]Asset, error)
	Get(id uint) (Asset, error)
	Edit(id uint, input Asset) (Asset, error)
	Delete(id uint) (error)
}
