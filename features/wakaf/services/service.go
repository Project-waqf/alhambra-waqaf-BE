package services

import (
	"wakaf/features/wakaf/domain"
	"wakaf/pkg/helper"
	paymentgateway "wakaf/utils/payment-gateway"

	"go.uber.org/zap"
)

type WakafService struct {
	WakafRepo domain.RepoInterface
}

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &WakafService{
		WakafRepo: data,
	}
}

var (
	logger = helper.Logger()
)

func (wakaf *WakafService) AddWakaf(input domain.Wakaf) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Insert(input)
	if err != nil {
		logger.Error("Failed insert wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) GetAllWakaf(category string, page int) ([]domain.Wakaf, int, error) {
	res, count, err := wakaf.WakafRepo.GetAllWakaf(category, page)
	if err != nil {
		logger.Error("Failed get all wakaf", zap.Error(err))
		return []domain.Wakaf{}, 0, err
	}
	return res, count, nil
}

func (wakaf *WakafService) UpdateWakaf(id uint, input domain.Wakaf) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Edit(id, input)
	if err != nil {
		logger.Error("Failed update wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) DeleteWakaf(id uint) (domain.Wakaf, error) {
	res, err := wakaf.WakafRepo.Delete(id)
	if err != nil {
		logger.Error("Failed delete wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) GetFileId(id uint) (string, error) {

	res, err := wakaf.WakafRepo.GetFileId(id)
	if err != nil {
		logger.Error("Failed get fileId", zap.Error(err))
		return res, err
	}
	return res, nil
}

func (wakaf *WakafService) GetSingleWakaf(id uint) (domain.Wakaf, error) {

	res, err := wakaf.WakafRepo.GetSingleWakaf(id)
	if err != nil {
		logger.Error("Failed get single wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) PayWakaf(input domain.PayWakaf) (domain.PayWakaf, error) {

	
	url, orderId := paymentgateway.PayBill(input)
	input.OrderId = orderId
	
	res, err := wakaf.WakafRepo.PayWakaf(input)
	if err != nil {
		logger.Error("Failed add donatur", zap.Error(err))
		return domain.PayWakaf{}, err
	}
	res.RedirectURL = url
	return res, nil
}

func (wakaf *WakafService) UpdatePayment(input domain.PayWakaf) (domain.PayWakaf, error) {

	res, err := wakaf.WakafRepo.UpdatePayment(input)
	if err != nil {
		logger.Error("Failed update payment", zap.Error(err))
		return domain.PayWakaf{}, err
	}
	
	return res, nil
}