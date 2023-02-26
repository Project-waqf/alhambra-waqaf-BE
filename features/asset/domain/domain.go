package domain

import "time"

type Asset struct {
	ID        uint
	Name      string
	Picture   string
	Detail    string
	Status    string
	FileId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsecaseInterface interface {
	AddAsset(input Asset) (Asset, error)
	GetAllAsset(status string, page int) ([]Asset, int, error)
	GetAsset(id uint ) (Asset, error)
	UpdateAsset(id uint, input Asset) (Asset, error)
	DeleteAsset(id uint) error
	ToOnline(id uint) error
	GetFileId(id uint) (string, error)
}

type RepositoryInterface interface {
	Insert(input Asset) (Asset, error)
	GetAll(status string, page int) ([]Asset, int, error)
	Get(id uint) (Asset, error)
	Edit(id uint, input Asset) (Asset, error)
	Delete(id uint) error
	ToOnline(id uint) error
	GetFileId(id uint) (string, error)
}
