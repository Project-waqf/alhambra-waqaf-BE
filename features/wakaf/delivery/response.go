package delivery

import (
	"wakaf/features/wakaf/domain"
)

type WakafResponse struct {
	ID        uint
	Title     string `json:"title"`
	Category  string `json:"category"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FromDomainAdd(input domain.Wakaf) WakafResponse {
	return WakafResponse{
		ID: input.ID,
		Title: input.Title,
		Category: input.Category,
		Picture: input.Picture,
		CreatedAt: input.CreatedAt.Format("Monday, 02-01-2006 T15:04:05"), 
		UpdatedAt: input.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05"),
	} 
}

func FromDomainGetAll(input []domain.Wakaf) []WakafResponse {
	var res []WakafResponse

	for _, v := range input {
		res = append(res, WakafResponse{
			ID: v.ID,
			Title: v.Title,
			Category: v.Category,
			Picture: v.Picture,
			CreatedAt:  v.CreatedAt.Format("Monday, 02-01-2006 T15:04:05"),
			UpdatedAt:  v.UpdatedAt.Format("Monday, 02-01-2006 T15:04:05"),
		})
	}
	return res
}