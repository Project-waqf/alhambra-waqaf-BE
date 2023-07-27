package services

import (
	"wakaf/features/news/domain"
	"wakaf/pkg/helper"

	"go.uber.org/zap"
)

type NewsServices struct {
	NewsRepository domain.RepoInterface
}

var (
	logger = helper.Logger()
)

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &NewsServices{
		NewsRepository: data,
	}
}

func (news *NewsServices) AddNews(input domain.News) (domain.News, error) {
	res, err := news.NewsRepository.Insert(input)
	if err != nil {
		logger.Error("Error insert data", zap.Error(err))
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) GetAll(status string, page int, sort string) ([]domain.News, int, int, int, error) {
	res, countOnline, countDraft, countArchive, err := news.NewsRepository.GetAll(status, page, sort)
	if err != nil {
		logger.Error("Error get all data", zap.Error(err))
		return nil, 0, 0, 0, err
	}
	return res, countOnline, countDraft, countArchive, nil
}

func (news *NewsServices) Get(id int) (domain.News, error) {
	res, err := news.NewsRepository.Get(id)
	if err != nil {
		logger.Error("Error get data", zap.Error(err))
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) UpdateNews(id int, input domain.News) (domain.News, error) {
	res, err := news.NewsRepository.Edit(id, input)
	if err != nil {
		logger.Error("Error update data", zap.Error(err))
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) Delete(id int) (domain.News, error) {
	res, err := news.NewsRepository.Delete(id)
	if err != nil {
		logger.Error("Error delete data", zap.Error(err))
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) ToOnline(id int) error {
	err := news.NewsRepository.ToOnline(id)
	if err != nil {
		logger.Error("Error to online data", zap.Error(err))
		return err
	}
	return nil
}

func (news *NewsServices) GetFileId(id int) (string, error) {

	res, err := news.NewsRepository.GetFileId(id) 
	if err != nil {
		logger.Error("Error get file id", zap.Error(err))
		return "", err
	}
	return res, nil
}

func (news *NewsServices) GetBySlug(slug string) (domain.News, error) {
	res, err := news.NewsRepository.GetBySlug(slug)
	if err != nil {
		logger.Error("Error get data", zap.Error(err))
		return domain.News{}, err
	}
	return res, nil
}