package services

import "wakaf/features/news/domain"

type NewsServices struct {
	NewsRepository domain.RepoInterface
}

func New(data domain.RepoInterface) domain.UseCaseInterface {
	return &NewsServices{
		NewsRepository: data,
	}
}

func (news *NewsServices) AddNews(input domain.News) (domain.News, error) {
	res, err := news.NewsRepository.Insert(input)
	if err != nil {
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) GetAll() ([]domain.News, error) {
	res, err := news.NewsRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (news *NewsServices) Get(id int) (domain.News, error) {
	res, err := news.NewsRepository.Get(id)
	if err != nil {
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) UpdateNews(id int, input domain.News) (domain.News, error) {
	res, err := news.NewsRepository.Edit(id, input)
	if err != nil {
		return domain.News{}, err
	}
	return res, nil
}

func (news *NewsServices) Delete(id int) error