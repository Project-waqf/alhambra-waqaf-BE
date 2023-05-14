package request

import "wakaf/features/partners/domain"

type PartnerRequest struct {
	Name    string `json:"name" form:"name"`
	Picture string `json:"picture" form:"name"`
	FileId  string
}

func ToDomainCreatePartner(input PartnerRequest) *domain.Partner {
	return &domain.Partner{
		Name:    input.Name,
		Picture: input.Picture,
		FileId:  input.FileId,
	}
}
