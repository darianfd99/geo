package repository

import localization "github.com/darianfd99/geo/pkg"

type LocalizationRepository interface {
	Save(loc localization.Localization) error
	GetAll() ([]localization.Localization, error)
}

//go:generate mockery --case=snake --outpkg=repositorymock --output=repositorymock --name=LocalizationRepository
type Repository struct {
	LocalizationRepository
}

func NewRepository() *Repository {
	return &Repository{
		LocalizationRepository: NewLocalizationSliceRepository(),
	}
}
