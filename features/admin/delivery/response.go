package delivery

import "wakaf/features/admin/domain"

type LoginResponse struct {
	ID       uint `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromDomainLogin(input domain.Admin) LoginResponse {
	return LoginResponse{
		ID:       input.ID,
		Name:     input.Name,
		Username: input.Username,
	}
}
