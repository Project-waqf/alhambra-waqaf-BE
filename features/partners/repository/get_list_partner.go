package repository

import (
	"wakaf/features/partners/domain"

	"gorm.io/gorm"
)

func (r *PartnerRepo) GetAll(limit, offset int, sort string) ([]*domain.Partner, error) {
	var res []Partner

	var query *gorm.DB
	if limit != 0 && offset != 0 {
		query = r.db.Limit(limit).Offset(offset).Order("created_at " + sort).Find(&res)
	} else if limit != 0 {
		query = r.db.Limit(limit).Order("created_at " + sort).Find(&res)
	} else if offset != 0 {
		query = r.db.Offset(offset).Order("created_at " + sort).Find(&res)
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return TodDomainGetListPartner(res), nil
}