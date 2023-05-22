package request

import "wakaf/features/partners/domain"

type PartnerRequest struct {
	Name        string `json:"name" form:"name"`
	PictureName string `json:"picture_name" form:"picture_name"`
	Link        string `json:"link" form:"string"`
	Picture     string
	FileId      string
}

func ToDomainCreatePartner(input PartnerRequest) *domain.Partner {
	return &domain.Partner{
		Name:        input.Name,
		PictureName: input.PictureName,
		Picture:     input.Picture,
		FileId:      input.FileId,
		Link:        input.Link,
	}
}
