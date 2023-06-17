package services

import "wakaf/features/asset/domain"

type AssetService struct {
	AssetRepo domain.RepositoryInterface
}

func New(data domain.RepositoryInterface) domain.UsecaseInterface {
	return &AssetService{
		AssetRepo: data,
	}
}

func (asset *AssetService) AddAsset(input domain.Asset) (domain.Asset, error) {
	res, err := asset.AssetRepo.Insert(input)
	if err != nil {
		return domain.Asset{}, err
	}
	return res, nil
}

func (asset *AssetService) GetAllAsset(status string, page int, sort string) ([]domain.Asset, int, int, int, error) {
	res, countOnline, countDraft, countArchive, err := asset.AssetRepo.GetAll(status, page, sort)
	if err != nil {
		return []domain.Asset{}, 0, 0, 0, err
	}
	return res, countOnline, countDraft, countArchive, nil
}

func (asset *AssetService) GetAsset(id uint) (domain.Asset, error) {
	res, err := asset.AssetRepo.Get(id)
	if err != nil {
		return domain.Asset{}, err
	}
	return res, nil
}

func (asset *AssetService) UpdateAsset(id uint, input domain.Asset) (domain.Asset, error) {
	res, err := asset.AssetRepo.Edit(id, input)
	if err != nil {
		return domain.Asset{}, err
	}
	return res, nil
}

func (asset *AssetService) DeleteAsset(id uint) error {
	err := asset.AssetRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (asset *AssetService) ToOnline(id uint) error {
	err := asset.AssetRepo.ToOnline(id)
	if err != nil {
		return err
	}
	return nil
}

func (asset *AssetService) GetFileId(id uint) (string, error) {
	res, err := asset.AssetRepo.GetFileId(uint(id))
	if err != nil {
		return "", err
	}
	return res, nil
}