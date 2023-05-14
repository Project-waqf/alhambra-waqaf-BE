package services

import (
	"wakaf/features/partners/domain"

	"go.uber.org/zap"
)

func (s *PartnersServices) UpdatePartner(id int, data *domain.Partner) (*domain.Partner, error) {
	res, err := s.PartnerRepo.Update(id, data)
	if err != nil {
		s.logger.Error("Failed update partner", zap.Error(err))
		return nil, err
	}
	return res, nil
}
