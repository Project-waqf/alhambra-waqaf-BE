package services

import (
	"wakaf/config"
	"wakaf/features/admin/domain"
	"wakaf/middlewares"

	"golang.org/x/crypto/bcrypt"
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

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(config.Getconfig().SALT1+input.Password+config.Getconfig().SALT2)); err != nil {
		return domain.Admin{}, err
	}

	token, err := middlewares.CreateToken(int(res.ID), res.Email)
	if err != nil {
		return domain.Admin{}, err
	}
	res.Token = token

	return res, nil
}

func (u *AdminServices) Register(input domain.Admin) error {

	if err := u.AdminRepository.GetUser(input); err != nil {
		return err
	}

	saltPw := config.Getconfig().SALT1 + input.Password + config.Getconfig().SALT2
	hash, err := bcrypt.GenerateFromPassword([]byte(saltPw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hash)
	err = u.AdminRepository.Register(input)
	if err != nil {
		return err
	}

	return nil
}

func (u *AdminServices) UpdatePassword(input domain.Admin) (error) {

	saltPw := config.Getconfig().SALT1 + input.Password + config.Getconfig().SALT2
	hash, err := bcrypt.GenerateFromPassword([]byte(saltPw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hash)
	if err := u.AdminRepository.UpdatePassword(input); err != nil {
		return err
	}
	return nil
}
