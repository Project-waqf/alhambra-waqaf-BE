package request

import "wakaf/features/partners/domain"

type PartnerRequest struct {
	Name        string `json:"name" form:"name"`
	Link        string `json:"link" form:"link"`
	Picture     string
	FileId      string
}

func ToDomainCreatePartner(input PartnerRequest) *domain.Partner {
	return &domain.Partner{
		Name:        input.Name,
		Picture:     input.Picture,
		FileId:      input.FileId,
		Link:        input.Link,
	}
}
