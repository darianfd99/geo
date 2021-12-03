package service

import (
	"errors"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/repository"
)

var ErrUuidExists = errors.New("the uuid received is already registered")

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

	exists := DoUuidExists(loc.Id, LocalizationsList)
	if exists {
		return *localization.ErrorGroupAvaerageLocation(), ErrUuidExists
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

func DoUuidExists(uuid localization.LocalizationUUID, locList []localization.Localization) bool {
	for _, loc := range locList {
		if loc.Id == uuid {
			return true
		}
	}
	return false
}
