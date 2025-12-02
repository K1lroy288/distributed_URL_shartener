package service

import (
	"shortener-service/model"
	"shortener-service/repository"
)

type ShortenerService struct {
	repo *repository.ShortenerRepository
}

func NewShortenerService(repo *repository.ShortenerRepository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) SaveCode(url *model.Url) (bool, error) {
	return s.repo.SaveCode(url)
}

func (s *ShortenerService) GetLink(code string) (*model.Url, error) {
	return s.repo.GetLink(code)
}
