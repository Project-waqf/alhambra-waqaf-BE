package response

import "wakaf/features/partners/domain"

type PartnerResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	PictureName string `json:"picture_name"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}

func GetListResponse(input []*domain.Partner) []PartnerResponse {
	var res []PartnerResponse

	for _, v := range input {
		tmp := GetDetailResponse(v)
		res = append(res, tmp)
	}
	return res
}

func GetDetailResponse(input *domain.Partner) PartnerResponse {
	return PartnerResponse{
		Id:          input.Id,
		Name:        input.Name,
		PictureName: input.PictureName,
		Picture:     input.Picture,
		CreatedAt:   input.CreatedAt.Format("Monday, 02-01-2006 T15:04:05"),
		UpdateAt:    input.UpdateAt.Format("Monday, 02-01-2006 T15:04:05"),
	}
}
