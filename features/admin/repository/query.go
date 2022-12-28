package repository

import (
	"wakaf/features/admin/domain"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.RepoInterface {
	return &AdminRepository{
		db: db,
	}
}

func (repo *AdminRepository) Login(data domain.Admin) (domain.Admin, error) {
	input := FromDomainLogin(data)
	
	if err := repo.db.Where("username = ? AND password = ?", input.Username, input.Password).First(&input).Error; err != nil {
		return domain.Admin{}, err
	}

	return ToDomainLogin(input), nil
}