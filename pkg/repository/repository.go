package repository

import localization "github.com/darianfd99/geo/pkg"

type LocalizationRepository interface {
	Save(loc localization.Localization) error
	GetAll() ([]localization.Localization, error)
}

type Repository struct {
	LocalizationRepository
}

func NewRepository() *Repository {
	return &Repository{
		LocalizationRepository: NewLocalizationSliceRepository(),
	}
}
