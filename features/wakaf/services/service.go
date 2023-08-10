package services

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"
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
	var emptyTime time.Time

	if input.DueDate == emptyTime {
		input.DueDate = time.Now()
	}

	res, err := wakaf.WakafRepo.Insert(input)
	if err != nil {
		logger.Error("Failed insert wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}

func (wakaf *WakafService) GetAllWakaf(category string, page int, isUser bool, sort, filter, status string) ([]domain.Wakaf, int, int, int, error) {
	res, countOnline, countDraft, countArchive, err := wakaf.WakafRepo.GetAllWakaf(category, page, isUser, sort, filter, status)
	if err != nil {
		logger.Error("Failed get all wakaf", zap.Error(err))
		return []domain.Wakaf{}, 0, 0, 0, err
	}
	return res, countOnline, countDraft, countArchive, nil
}

func isEmptyValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return value.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0.0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Ptr, reflect.Interface:
		return value.IsNil()
	default:
		return false
	}
}

func (wakaf *WakafService) UpdateWakaf(id uint, input domain.Wakaf) (domain.Wakaf, error) {

	resGet, err := wakaf.WakafRepo.GetSingleWakaf(id)
	if err != nil {
		logger.Error("Wakaf not found", zap.Error(err))
		return domain.Wakaf{}, err
	}

	res, err := wakaf.WakafRepo.Edit(id, resGet)
	if err != nil {
		logger.Error("Failed update wakaf", zap.Error(err))
		return res, err
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

	snapResp, orderId, err := paymentgateway.Midtrans(input)
	if err != nil {
		logger.Info("Failed payment gateway", zap.Error(err))
		return domain.PayWakaf{}, err
	}
	input.OrderId = orderId
	input.Status = "pending"
	if snapResp.RedirectURL == "" {
		logger.Info("Failed get redirect url because transaction has completed")
		return domain.PayWakaf{}, errors.New("completed")
	}

	err = wakaf.WakafRepo.SaveRedis(orderId, input)
	if err != nil {
		logger.Error("Failed save donor to redis", zap.Error(err))
		return domain.PayWakaf{}, err
	}

	input.RedirectURL = snapResp.RedirectURL
	input.Token = snapResp.Token
	return input, nil
}

func (wakaf *WakafService) UpdatePayment(input domain.PayWakaf) (domain.PayWakaf, error) {

	resRedis, err := wakaf.WakafRepo.GetFromRedis(input.OrderId)
	if err != nil {
		logger.Error("Failed get data donor from redis")
		return domain.PayWakaf{}, err
	}

	var dataDonor domain.PayWakaf
	if err := json.Unmarshal([]byte(resRedis), &dataDonor); err != nil {
		logger.Error("Error unmarshal data donor", zap.Error(err))
		return domain.PayWakaf{}, err
	}

	_, err = wakaf.WakafRepo.UpdatePayment(dataDonor)
	if err != nil {
		logger.Error("Failed update payment", zap.Error(err))
		return domain.PayWakaf{}, err
	}

	resDonor, err := wakaf.WakafRepo.PayWakaf(dataDonor)
	if err != nil {
		logger.Error("Failed insert donor", zap.Error(err))
		return domain.PayWakaf{}, err
	}

	return resDonor, nil
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

func (wakaf *WakafService) SearchWakaf(input string) ([]domain.Wakaf, int, int, int, error) {

	res, countOnline, countDraft, countArchive, err := wakaf.WakafRepo.Search(input)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	return res, countOnline, countDraft, countArchive, nil
}

func (wakaf *WakafService) GetSummary() (int, int, int, error) {

	count, sum, wakif, err := wakaf.WakafRepo.GetSummary()
	if err != nil {
		logger.Error("Failed get summary wakaf", zap.Error(err))
		return 0, 0, 0, err
	}
	return count, sum, wakif, nil
}

func (wakaf *WakafService) GetSummaryDashboard() (int, int, int, error) {

	online, complete, asset, err := wakaf.WakafRepo.GetSummaryDashboard()
	if err != nil {
		logger.Error("Failed get summary wakaf dashboard")
		return 0, 0, 0, err
	}
	return online, complete, asset, nil
}

func (wakaf *WakafService) GetSingleWakafBySlug(slug string) (domain.Wakaf, error) {

	res, err := wakaf.WakafRepo.GetSingleWakafBySlug(slug)
	if err != nil {
		logger.Error("Failed get single wakaf", zap.Error(err))
		return domain.Wakaf{}, err
	}
	return res, nil
}
