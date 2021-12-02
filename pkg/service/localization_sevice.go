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

func (s *LocalizationService) Post(id string, lat, long float64) (localization.GroupAvaerageLocation, error) {
	loc, err := localization.NewLocalization(id, lat, long)
	if err != nil {
		return *localization.ErrorGroupAvaerageLocation(), err
	}

	LocalizationsList, err := s.repo.GetAll()
	if err != nil {
		return *localization.ErrorGroupAvaerageLocation(), err
	}

	err = s.repo.Save(*loc)
	if err != nil {
		return *localization.ErrorGroupAvaerageLocation(), err
	}

	group := loc.GetGroupAvaerageLocation(LocalizationsList)
	return *group, nil

}

func (s *LocalizationService) GetAll() ([]localization.Localization, error) {
	LocalizationsList, err := s.repo.GetAll()
	return LocalizationsList, err
}
