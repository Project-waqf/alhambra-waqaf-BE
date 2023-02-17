package repository

import (
	"errors"
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

	if err := repo.db.Where("email = ?", input.Email).First(&input).Error; err != nil {
		return domain.Admin{}, err
	}

	return ToDomainLogin(input), nil
}

func (repo *AdminRepository) Register(data domain.Admin) error {
	input := FromDomainRegister(data)

	if err := repo.db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepository) GetUser(data domain.Admin) error {
	var res Admin

	if err := repo.db.Where("email", data.Email).First(&res).Error; err == nil {
		return errors.New("email has taken")
	}
	return nil
}
