package domain

import "time"

type Partner struct {
	Id          uint
	Name        string
	Picture     string
	FileId      string
	Link        string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type UseCaseInterface interface {
	GetAllPartner(limit, offset int, sort string) ([]*Partner, error)
	GetPartnerDetail(id int) (*Partner, error)
	UpdatePartner(id int, data *Partner) (*Partner, error)
	CreatePartner(data *Partner) (*Partner, error)
	DeletePartner(id int) error
	GetFileId(id int) (string, error)
}

type RepoInterface interface {
	GetAll(limit, offset int, sort string) ([]*Partner, error)
	GetDetail(id int) (*Partner, error)
	Update(id int, data *Partner) (*Partner, error)
	Insert(data *Partner) (*Partner, error)
	Delete(id int) error
	GetFileId(id int) (string, error)
}
