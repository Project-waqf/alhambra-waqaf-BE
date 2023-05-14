package services

import (
	"wakaf/features/partners/domain"

	"go.uber.org/zap"
)

func (s *PartnersServices) GetPartnerDetail(id int) (*domain.Partner, error) {
	res, err := s.PartnerRepo.GetDetail(id)
	if err != nil {
		s.logger.Error("Failed get detail partner", zap.Error(err))
		return nil, err
	}
	return res, nil
}