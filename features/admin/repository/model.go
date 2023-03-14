package repository

import (
	"wakaf/features/admin/domain"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func FromDomainLogin(input domain.Admin) Admin {
	return Admin{
		Email:    input.Email,
		Password: input.Password,
	}
}

func FromDomainRegister(input domain.Admin) Admin {
	return Admin{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

func ToDomainLogin(input Admin) domain.Admin {
	return domain.Admin{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}
