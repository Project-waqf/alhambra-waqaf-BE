package services

import (
	"wakaf/features/partners/domain"

	"go.uber.org/zap"
)

func (s *PartnersServices) GetAllPartner(limit, offset int, sort string) ([]*domain.Partner, error) {
	res, err := s.PartnerRepo.GetAll(limit, offset, sort)
	if err != nil {
		s.logger.Error("Failed get all partner", zap.Error(err))
		return nil, err
	}
	return res, nil
}
