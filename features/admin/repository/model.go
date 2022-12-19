package repository

import (
	"wakaf/features/admin/domain"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Username string
	Password string
}

func FromDomainLogin(input domain.Admin) Admin {
	return Admin{
		Username: input.Username,
		Password: input.Password,
	}
}

func ToDomainLogin(input Admin) domain.Admin {
	return domain.Admin{
		ID: input.ID,
		Name: input.Name,
		Username: input.Username,
	}
}