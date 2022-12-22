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