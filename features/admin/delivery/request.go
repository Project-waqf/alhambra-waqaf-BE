package delivery

import "wakaf/features/admin/domain"

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Register struct {
	Id       int
	Name     string `json:"name,omitempty" form:"name"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	OldPassword string `json:"old_password,omitempty" form:"old_password"`
}

type RegisterResponseNew struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ForgotUpdate struct {
	Token    string `json:"token" form:"token"`
	Password string `json:"password" form:"password"`
}

type Forgot struct {
	Email string `json:"email" form:"email"`
}

func ToDomainLogin(data Login) domain.Admin {
	return domain.Admin{
		Email:    data.Email,
		Password: data.Password,
	}
}

func ToDomainRegister(data Register) domain.Admin {
	return domain.Admin{
		ID:       uint(data.Id),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
