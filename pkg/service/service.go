package service

import (
	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/repository"
)

type Localization interface {
	Post(id string, lat float64, long float64) (localization.GroupAvaerageLocation, error)
	GetAll() ([]localization.Localization, error)
}

//go:generate mockery --case=snake --outpkg=servicemocks --output=servicemocks --name=Localization

type Service struct {
	Localization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Localization: NewLocalizationService(repos.LocalizationRepository),
	}
}
