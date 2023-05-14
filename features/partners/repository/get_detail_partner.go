package repository

import "wakaf/features/partners/domain"

func (r *PartnerRepo) GetDetail(id int) (*domain.Partner, error) {
	var data Partner

	err := r.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return ToDomainCreatePartner(data), nil
}