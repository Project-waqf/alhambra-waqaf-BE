package delivery

import "wakaf/features/asset/domain"

type AssetResponse struct {
	ID        uint	`json:"id_asset"`
	Name      string `json:"name"`
	Picture   string `json:"picture""`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomainAdd(input domain.Asset) AssetResponse {
	return AssetResponse{
		ID: input.ID,
		Name: input.Name,
		Picture: input.Picture,
		Detail: input.Detail,
		CreatedAt: input.CreatedAt.Format("Monday, 02-01-2006 T15:04:05"),
		UpdatedAt: input.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05"),
	}
}