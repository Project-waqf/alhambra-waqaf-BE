package delivery

import "wakaf/features/admin/domain"

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func ToDomainLogin(data Login) domain.Admin {
	return domain.Admin{
		Username: data.Username,
		Password: data.Password,
	}
}