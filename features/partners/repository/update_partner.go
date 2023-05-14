package repository

import "wakaf/features/partners/domain"

func (r *PartnerRepo) Update(id int, input *domain.Partner) (*domain.Partner, error) {
	data := FromDomainCreatePartner(input)

	err := r.db.Where("id = ?", id).Updates(&data).Last(&data).Error
	if err != nil {
		return nil, err
	}
	return ToDomainCreatePartner(data), nil
}