package delivery

import "wakaf/features/asset/domain"

type AssetRequest struct {
	ID      uint
	Name    string `json:"name" form:"name"`
	Picture string `json:"picture" form:"picture"`
	Detail  string `json:"detail" form:"detail"`
}

func ToDomainAdd(input AssetRequest) domain.Asset {
	return domain.Asset{
		ID: input.ID,
		Name: input.Name,
		Picture: input.Picture,
		Detail: input.Detail,
	}
}
