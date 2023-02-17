package delivery

import "wakaf/features/admin/domain"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromDomainLogin(input domain.Admin) LoginResponse {
	return LoginResponse{
		ID:   input.ID,
		Name: input.Name,
		Email:   input.Email,
		Token: input.Token,
	}
}
