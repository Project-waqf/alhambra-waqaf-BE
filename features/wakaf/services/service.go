package services

import "wakaf/features/wakaf/domain"

type WakafService struct {
	WakafRepo domain.RepoInterface
}

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &WakafService{
		WakafRepo: data,
	}
}

func (wakaf *WakafService) AddWakaf(input domain.Wakaf) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Insert(input)
	if err != nil {
		return domain.Wakaf{}, err
	}
	return res, nil
}

func(wakaf *WakafService) GetAllWakaf() ([]domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.GetAllWakaf()
	if err != nil {
		return []domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) UpdateWakaf(id uint, input domain.Wakaf) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Edit(id, input)
	if err != nil {
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) DeleteWakaf(id uint) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Delete(id)
	if err != nil {
		return domain.Wakaf{}, err
	}
	return res, nil
}
