package repository

import "wakaf/features/partners/domain"

func (r *PartnerRepo) Insert(input *domain.Partner) (*domain.Partner, error) {
	data := FromDomainCreatePartner(input)
	err := r.db.Create(&data).Last(&data).Error
	if err != nil {
		return nil, err
	}
	return ToDomainCreatePartner(data), nil
}