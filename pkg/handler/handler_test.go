package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/service"
	"github.com/darianfd99/geo/pkg/service/servicemocks"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerPostLocalization(t *testing.T) {
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	lat := 1.0
	long := 1.0

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(localization.GroupAvaerageLocation{}, nil)
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()
		postLocalizationsRequst := localizationRequest{
			Id:  "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Lat: 2.34,
		}

		b, err := json.Marshal(postLocalizationsRequst)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertNotCalled(t, mock.Anything)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given an invalid uuid it returns 400", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("Post", "invalid uuid", lat, long).Return(localization.GroupAvaerageLocation{}, localization.ErrInvalidCourseID)
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()

		postLocalizationsRequst := localizationRequest{
			Id:   "invalid uuid",
			Lat:  lat,
			Long: long,
		}

		b, err := json.Marshal(postLocalizationsRequst)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertExpectations(t)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given an unexpected error from service layer it returns 500", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(localization.GroupAvaerageLocation{}, errors.New("unexected error"))
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()
		postLocalizationsRequst := localizationRequest{
			Id:   id,
			Lat:  lat,
			Long: long,
		}

		b, err := json.Marshal(postLocalizationsRequst)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(localization.GroupAvaerageLocation{}, nil)
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()

		postLocalizationsRequst := localizationRequest{
			Id:   "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Lat:  2.34,
			Long: 1,
		}

		b, err := json.Marshal(postLocalizationsRequst)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertExpectations(t)
		assert.Equal(t, LocationService.AssertExpectations(t), true)
		assert.Equal(t, http.StatusCreated, res.StatusCode)

	})
}

func TestHandlerGetAllLocalizations(t *testing.T) {
	id := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	lat := 1.0
	long := 1.0

	t.Run("given an unexpected error from service layer it returns 500", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("GetAll").Return([]localization.Localization{}, errors.New("unexected error"))
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()
		postLocalizationsRequst := localizationRequest{
			Id:   id,
			Lat:  lat,
			Long: long,
		}

		b, err := json.Marshal(postLocalizationsRequst)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	})

	t.Run("all ok it returns 200", func(t *testing.T) {
		LocationService := new(servicemocks.Localization)
		LocationService.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return([]localization.Localization{}, nil)
		service := &service.Service{Localization: LocationService}

		gin.SetMode(gin.TestMode)
		handler := NewHandler(service)
		r := handler.InitRoutes()

		b, err := json.Marshal(map[string]interface{}{})
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		LocationService.AssertExpectations(t)
		assert.Equal(t, LocationService.AssertExpectations(t), true)
		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

}
