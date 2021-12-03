package service

import (
	"errors"
	"testing"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/repository"
	"github.com/darianfd99/geo/pkg/repository/repositorymock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLocalizationServicePost(t *testing.T) {
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	lat := 1.0
	long := 1.0

	loc, err := localization.NewLocalization(id, lat, long)
	require.NoError(t, err)

	t.Run("error in repository", func(t *testing.T) {
		mockRepository := new(repositorymock.LocalizationRepository)
		mockRepository.On("Save", *loc).Return(errors.New("something unexpexted"))
		mockRepository.On("GetAll").Return([]localization.Localization{}, nil)
		repos := &repository.Repository{
			LocalizationRepository: mockRepository,
		}

		services := NewService(repos)

		_, err = services.Localization.Post(id, lat, long)
		mockRepository.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("invalid uuid", func(t *testing.T) {
		mockRepository := new(repositorymock.LocalizationRepository)
		mockRepository.On("Save", *loc).Return(nil)
		mockRepository.On("GetAll").Return([]localization.Localization{}, nil)

		repos := &repository.Repository{
			LocalizationRepository: mockRepository,
		}

		services := NewService(repos)

		_, err = services.Localization.Post("invalidUUID", lat, long)
		mockRepository.AssertNotCalled(t, mock.Anything)
		assert.Error(t, err)
		assert.Equal(t, errors.Is(err, localization.ErrInvalidCourseID), true)

	})

	t.Run("uuid repeated", func(t *testing.T) {
		mockRepository := new(repositorymock.LocalizationRepository)
		mockRepository.On("Save", mock.Anything).Return(nil)
		mockRepository.On("GetAll").Return([]localization.Localization{*loc}, nil)

		repos := &repository.Repository{
			LocalizationRepository: mockRepository,
		}

		services := NewService(repos)

		_, err = services.Localization.Post(loc.Id.String(), lat, long)
		mockRepository.AssertCalled(t, "GetAll")
		mockRepository.AssertNotCalled(t, "Save")

		assert.Error(t, err)
		assert.Equal(t, errors.Is(err, ErrUuidExists), true)
	})

	t.Run("3 geolocations less than one kilometer and one more than one kilometer", func(t *testing.T) {
		loc1, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723a", 1.1, 1.1)
		require.NoError(t, err)
		loc2, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723e", 1.2, 1.2)
		require.NoError(t, err)
		loc3, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723b", 145.2, 154.2)
		require.NoError(t, err)

		expectedGroup1 := *localization.NewGroupAvaerageLocation(id, lat, long)
		expectedGroup2 := *localization.NewGroupAvaerageLocation(loc1.Id.String(), (lat+loc1.Lat)/2, (long+loc1.Long)/2)
		expectedGroup3 := *localization.NewGroupAvaerageLocation(loc2.Id.String(), (lat+loc1.Lat+loc2.Lat)/3, (long+loc1.Long+loc2.Long)/3)
		expectedGroup4 := *localization.NewGroupAvaerageLocation(loc3.Id.String(), loc3.Lat, loc3.Long)

		repos := &repository.Repository{
			LocalizationRepository: repository.NewLocalizationSliceRepository(),
		}

		services := NewService(repos)

		group1, err := services.Localization.Post(id, lat, long)
		assert.NoError(t, err)
		assert.Equal(t, expectedGroup1, group1, "Exepected: %v, received: %v", expectedGroup1, group1)

		group2, err := services.Localization.Post(loc1.Id.String(), loc1.Lat, loc1.Long)
		assert.NoError(t, err)
		assert.Equal(t, expectedGroup2, group2, "Exepected: %v, received: %v", expectedGroup1, group1)

		group3, err := services.Localization.Post(loc2.Id.String(), loc2.Lat, loc2.Long)
		assert.NoError(t, err)
		assert.Equal(t, expectedGroup3, group3, "Exepected: %v, received: %v", expectedGroup1, group1)

		group4, err := services.Localization.Post(loc3.Id.String(), loc3.Lat, loc3.Long)
		assert.NoError(t, err)
		assert.Equal(t, expectedGroup4, group4, "Exepected: %v, received: %v", expectedGroup1, group1)
	})

}

func TestLocalizationServiceGetAll(t *testing.T) {
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	lat := 1.0
	long := 1.0

	loc, err := localization.NewLocalization(id, lat, long)
	require.NoError(t, err)

	t.Run("error in repository", func(t *testing.T) {
		mockRepository := new(repositorymock.LocalizationRepository)
		mockRepository.On("GetAll").Return([]localization.Localization{}, errors.New("something unexpexted"))
		repos := &repository.Repository{
			LocalizationRepository: mockRepository,
		}

		services := NewService(repos)

		_, err = services.Localization.Post(id, lat, long)
		mockRepository.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("3 geolocations less than one kilometer and one more than one kilometer", func(t *testing.T) {
		loc1, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723a", 1.1, 1.1)
		require.NoError(t, err)
		loc2, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723e", 1.2, 1.2)
		require.NoError(t, err)
		loc3, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723b", 145.2, 154.2)
		require.NoError(t, err)

		expectedList := []localization.Localization{*loc, *loc1, *loc2, *loc3}
		repos := &repository.Repository{
			LocalizationRepository: repository.NewLocalizationSliceRepository(),
		}

		services := NewService(repos)

		_, err = services.Localization.Post(id, lat, long)
		require.NoError(t, err)
		_, err = services.Localization.Post(loc1.Id.String(), loc1.Lat, loc1.Long)
		require.NoError(t, err)
		_, err = services.Localization.Post(loc2.Id.String(), loc2.Lat, loc2.Long)
		require.NoError(t, err)
		_, err = services.Localization.Post(loc3.Id.String(), loc3.Lat, loc3.Long)
		require.NoError(t, err)

		list, err := services.Localization.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, expectedList, list, "Exepected: %v, received: %v", expectedList, list)
	})

}
