package delivery

import "wakaf/features/admin/domain"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Image string `json:"image"`
}

type ProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

func FromDomainLogin(input domain.Admin) LoginResponse {
	return LoginResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Token: input.Token,
		Image: input.Image,
	}
}

func FromDomainProfile(input domain.Admin) ProfileResponse {
	return ProfileResponse{
		Name:  input.Name,
		Email: input.Email,
		Image: input.Image,
	}
}
