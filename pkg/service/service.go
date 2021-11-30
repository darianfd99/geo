package service

import (
	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/repository"
)

type Localization interface {
	Post(loc localization.Localization) error
	GetAll() ([]localization.Localization, error)
}

type Service struct {
	Localization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Localization: NewLocalizationService(repos.LocalizationRepository),
	}
}
