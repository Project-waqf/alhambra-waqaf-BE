package delivery

import "wakaf/features/admin/domain"

type LoginResponse struct {
	ID       uint   
	Name     string 
	Username string
	Password string 
}

func FromDomainLogin(input domain.Admin) LoginResponse {
	return LoginResponse{
		ID: input.ID,
		Name: input.Name,
		Username: input.Username,
		Password: input.Password,
	}
}
