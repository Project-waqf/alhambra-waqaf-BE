package delivery

import "wakaf/features/asset/domain"

type AssetRequest struct {
	ID      uint
	Name    string `json:"name" form:"name"`
	Picture string `json:"picture" form:"picture"`
	Detail  string `json:"detail" form:"detail"`
	Type    string `json:"type" form:"type"`
	FileId  string
}

func ToDomainAdd(input AssetRequest) domain.Asset {
	return domain.Asset{
		ID:      input.ID,
		Name:    input.Name,
		Picture: input.Picture,
		Detail:  input.Detail,
		Type:    input.Type,
		FileId:  input.FileId,
	}
}
