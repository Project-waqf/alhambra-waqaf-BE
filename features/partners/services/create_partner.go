package services

import (
	"wakaf/features/partners/domain"

	"go.uber.org/zap"
)

func (s *PartnersServices) CreatePartner(input *domain.Partner) (*domain.Partner, error) {
	res, err := s.PartnerRepo.Insert(input)
	if err != nil {
		s.logger.Error("Failed create new partner", zap.Error(err))
		return nil, err
	}
	return res, nil
}