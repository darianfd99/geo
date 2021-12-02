package repository

import (
	"testing"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocalizationSliceRepositorySave(t *testing.T) {
	repos := NewRepository()
	repo := repos.LocalizationRepository

	loc1, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723f", 2.34, 1)
	require.NoError(t, err)
	loc2, err := localization.NewLocalization("8a1c5cdc-ba57-445a-994d-aa412d23723a", 23.4, 22)
	require.NoError(t, err)

	expectedLocsList := []localization.Localization{*loc1, *loc2}

	t.Run("Saving 2 localization objects in the slice and getting all objects", func(t *testing.T) {
		err = repo.Save(*loc1)
		assert.NoError(t, err)
		err = repo.Save((*loc2))
		assert.NoError(t, err)

		locList, err := repo.GetAll()
		assert.NoError(t, err)
		assert.EqualValues(t, expectedLocsList, locList)

	})

}
