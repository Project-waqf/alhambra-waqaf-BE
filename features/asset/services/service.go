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