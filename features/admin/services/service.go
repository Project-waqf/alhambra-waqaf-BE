package services

import (
	"wakaf/features/admin/domain"
)

type AdminServices struct {
	AdminRepository domain.RepoInterface
}

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &AdminServices{
		AdminRepository: data,
	}
}

func (service *AdminServices) Login(input domain.Admin) (domain.Admin, error) {
	res, err := service.AdminRepository.Login(input)
	if err != nil {
		return domain.Admin{}, err
	}
	return res, nil
}