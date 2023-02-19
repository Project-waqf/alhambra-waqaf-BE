package delivery

import "wakaf/features/wakaf/domain"

type WakafRequest struct {
	ID       uint
	Title    string `json:"title" form:"title"`
	Category string `json:"category" form:"category"`
	Picture  string `json:"picture" form:"picture"`
	FileId   string
}

func ToDomainAdd(input WakafRequest) domain.Wakaf {
	return domain.Wakaf{
		ID:       input.ID,
		Title:    input.Title,
		Category: input.Category,
		Picture:  input.Picture,
		FileId:   input.FileId,
	}
}
