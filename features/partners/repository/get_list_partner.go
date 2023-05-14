package repository

import (
	"wakaf/features/partners/domain"

	"gorm.io/gorm"
)

func (r *PartnerRepo) GetAll(limit, offset int) ([]*domain.Partner, error) {
	var res []Partner

	var query *gorm.DB
	if limit != 0 && offset != 0 {
		query = r.db.Limit(limit).Offset(offset).Find(&res)
	} else if limit != 0 {
		query = r.db.Limit(limit).Find(&res)
	} else if offset != 0 {
		query = r.db.Offset(offset).Find(&res)
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return TodDomainGetListPartner(res), nil
}