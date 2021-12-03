package benchmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	localization "github.com/darianfd99/geo/pkg"
	"github.com/darianfd99/geo/pkg/handler"
	"github.com/darianfd99/geo/pkg/repository"
	"github.com/darianfd99/geo/pkg/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func BenchmarkPostWithEmptyListLocations(b *testing.B) {
	id, err := uuid.NewV4()
	require.NoError(b, err)
	lat := 1.0
	long := 1.0

	postLocalizationsRequst := handler.LocalizationRequest{
		Id:   id.String(),
		Lat:  &lat,
		Long: &long,
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)

	gin.SetMode(gin.TestMode)
	handlers := handler.NewHandler(service)
	r := handlers.InitRoutes()

	body, err := json.Marshal(postLocalizationsRequst)
	require.NoError(b, err)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		r.ServeHTTP(rec, req)
		b.StopTimer()

		_, _ = repos.DeleteAll()
		list, _ := repos.LocalizationRepository.GetAll()
		fmt.Printf("len: %d ", len(list))
		id, err = uuid.NewV4()
		require.NoError(b, err)
		lat++
		long++
		postLocalizationsRequst = handler.LocalizationRequest{
			Id:   id.String(),
			Lat:  &lat,
			Long: &long,
		}

		body, err := json.Marshal(postLocalizationsRequst)
		require.NoError(b, err)

		rec = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	}

}

func BenchmarkPostWithListLocations(b *testing.B) {
	lat := 1.0
	long := 1.0
	id, _ := uuid.NewV4()
	list := []localization.Localization{}
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		loc, err := localization.NewLocalization(id.String(), lat, long)
		require.NoError(b, err)
		list = append(list, *loc)

		id, _ = uuid.NewV4()
		lat = lat + 0.1
		long = long + 0.1
	}

	id, err := uuid.NewV4()
	require.NoError(b, err)
	lat = 1.0
	long = 1.0

	postLocalizationsRequst := handler.LocalizationRequest{
		Id:   id.String(),
		Lat:  &lat,
		Long: &long,
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)

	gin.SetMode(gin.TestMode)
	handlers := handler.NewHandler(service)
	r := handlers.InitRoutes()

	body, err := json.Marshal(postLocalizationsRequst)
	require.NoError(b, err)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		r.ServeHTTP(rec, req)
		b.StopTimer()

		_ = repos.LocalizationRepository.SetList(list)
		list, _ := repos.LocalizationRepository.GetAll()
		fmt.Printf("len: %d ", len(list))
		id, err = uuid.NewV4()
		require.NoError(b, err)
		lat = lat + 0.01
		long = lat + 0.01
		postLocalizationsRequst = handler.LocalizationRequest{
			Id:   id.String(),
			Lat:  &lat,
			Long: &long,
		}

		body, err := json.Marshal(postLocalizationsRequst)
		require.NoError(b, err)

		rec = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	}

}

func BenchmarkGetAll(b *testing.B) {
	repos := repository.NewRepository()
	service := service.NewService(repos)

	gin.SetMode(gin.TestMode)
	handler := handler.NewHandler(service)
	r := handler.InitRoutes()

	postLocalizationsRequst := struct{}{}

	body, err := json.Marshal(postLocalizationsRequst)
	require.NoError(b, err)

	req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
	require.NoError(b, err)

	rec := httptest.NewRecorder()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.ServeHTTP(rec, req)
	}

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(b, http.StatusOK, res.StatusCode)
}
