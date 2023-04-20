package services

import (
	"errors"
	"strings"
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

	resWakaf, err := wakaf.WakafRepo.GetSingleWakaf(uint(input.IdWakaf))
	if err != nil {
		return domain.PayWakaf{}, err
	}

	if (resWakaf.Collected + input.GrossAmount) > resWakaf.FundTarget {
		input.GrossAmount = resWakaf.FundTarget - resWakaf.Collected
	}

	url, orderId := paymentgateway.PayBill(input)
	input.OrderId = orderId
	input.Status = "pending"
	if url == "" {
		return domain.PayWakaf{}, errors.New("completed")
	}
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

func (wakaf *WakafService) DenyTransaction(input string) error {

	res, err := paymentgateway.DenyTransaction(input)
	if err != nil {
		return err
	}

	if !strings.Contains(res, "200") {
		var err = errors.New("failed to deny transacton")
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (wakaf *WakafService) SearchWakaf(input string) ([]domain.Wakaf, int, error) {

	res, err := wakaf.WakafRepo.Search(input)
	if err != nil {
		return nil, 0, err
	}
	return res, len(res), nil
}

func (wakaf *WakafService) GetSummary() (int, int, int, error) {
	
	count, sum, wakif, err := wakaf.WakafRepo.GetSummary()
	if err != nil {
		logger.Error("Failed get summary wakaf", zap.Error(err))
		return 0, 0, 0, err
	}
	return count, sum, wakif, nil
}