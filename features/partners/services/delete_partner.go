package services

import "go.uber.org/zap"

func (s *PartnersServices) DeletePartner(id int) error {
	err := s.PartnerRepo.Delete(id)
	if err != nil {
		s.logger.Error("Failed delete partner", zap.Error(err))
		return err
	}
	return err
}
