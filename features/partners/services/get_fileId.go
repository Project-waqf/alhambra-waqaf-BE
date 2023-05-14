package services

func (s *PartnersServices) GetFileId(id int) (string, error) {
	res, err := s.PartnerRepo.GetFileId(id)
	if err != nil {
		return "", err
	}
	return res, nil
}