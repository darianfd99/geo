package service

import (
	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/repository"
)

type LocalizationService struct {
	repo repository.LocalizationRepository
}

func NewLocalizationService(repo repository.LocalizationRepository) *LocalizationService {
	return &LocalizationService{
		repo: repo,
	}
}

func (s *LocalizationService) Post(loc localization.Localization) error {
	error := s.repo.Save(loc)
	return error
}

func (s *LocalizationService) GetAll() ([]localization.Localization, error) {
	LocalizationsList, error := s.repo.GetAll()
	return LocalizationsList, error
}
