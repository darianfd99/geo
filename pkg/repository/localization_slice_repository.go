package repository

import localization "github.com/darianfd99/geo/pkg"

type LocalizationSliceRepository struct {
	LocalizationsSlice []localization.Localization
}

func NewLocalizationSliceRepository() *LocalizationSliceRepository {
	return &LocalizationSliceRepository{
		LocalizationsSlice: []localization.Localization{},
	}
}

func (r *LocalizationSliceRepository) Save(loc localization.Localization) error {

	r.LocalizationsSlice = append(r.LocalizationsSlice, loc)
	return nil
}

func (r *LocalizationSliceRepository) GetAll() ([]localization.Localization, error) {
	return r.LocalizationsSlice, nil
}

func (r *LocalizationSliceRepository) DeleteAll() ([]localization.Localization, error) {
	r.LocalizationsSlice = []localization.Localization{}
	return r.LocalizationsSlice, nil
}

func (r *LocalizationSliceRepository) SetList(list []localization.Localization) error {
	r.LocalizationsSlice = list
	return nil
}
