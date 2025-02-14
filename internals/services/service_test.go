package services_test

import (
	"sync"
	"testing"

	"github.com/Akshayvij07/country-search/pkg/errors"

	"github.com/Akshayvij07/country-search/internals/mocks"
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/Akshayvij07/country-search/internals/services"
	"github.com/stretchr/testify/assert"
)

// Mock third-party API
func mockCountryResponse() *models.CountryResponse {
	return &models.CountryResponse{
		Name:    models.Name{Common: "India"},
		Capital: []string{"New Delhi"},
		Currencies: map[string]models.Currency{
			"INR": {Name: "Indian Rupee",
				Symbol: "₹",
			}},
		Population: 1380004385,
	}
}

func TestGetCountry_FromCache(t *testing.T) {
	mockCache := new(mocks.MockCache)
	mockFetcher := new(mocks.MockFetcher)
	mockService := services.New(mockCache, mockFetcher)

	expectedCountry := models.NewCountry("India", "New Delhi", "INR", 1380004385)

	// Cache hit
	mockCache.On("Get", "India").Return(expectedCountry, nil)

	country, err := mockService.GetCountry("India")
	assert.NoError(t, err)
	assert.Equal(t, expectedCountry, country)

	mockCache.AssertExpectations(t)
}

func TestGetCountry_FromThirdPartyAPI(t *testing.T) {
	mockCache := new(mocks.MockCache)
	mockFetcher := new(mocks.MockFetcher)
	mockService := services.New(mockCache, mockFetcher)

	expectedCountry := models.NewCountry("India", "New Delhi", "INR", 1380004385)

	// Cache miss
	mockCache.On("Get", "India").Return(nil, errors.ErrKeyNotFound)
	mockCache.On("Set", "India", expectedCountry).Return()

	// Mock third-party API response
	mockFetcher.On("FetchCountry", "India").Return(mockCountryResponse(), nil)

	country, err := mockService.GetCountry("India")
	assert.NoError(t, err)
	assert.Equal(t, expectedCountry, country)

	mockCache.AssertExpectations(t)
}

func TestGetCountry_InvalidName(t *testing.T) {
	mockCache := new(mocks.MockCache)
	mockFetcher := new(mocks.MockFetcher)
	mockService := services.New(mockCache, mockFetcher)

	// Invalid country name
	mockFetcher.On("FetchCountry", "Invalid").Return(mockCountryResponse(), nil)
	mockCache.On("Get", "Invalid").Return(nil, errors.ErrKeyNotFound)

	_, err := mockService.GetCountry("Invalid")
	assert.Error(t, err)
	assert.Equal(t, "key not found", err.Error())
}

// RaceConditionTest ConcurrentAPIRequests tests concurrent requests to the API
func TestConcurrentAPIRequests(t *testing.T) {
	mockCache := new(mocks.MockCache)
	mockFetcher := new(mocks.MockFetcher)
	service := services.New(mockCache, mockFetcher)

	expectedCountry := models.NewCountry("India", "New Delhi", "INR", 1380004385)

	mockCache.On("Get", "India").Return(nil, errors.ErrKeyNotFound)
	mockCache.On("Set", "India", expectedCountry).Return()

	mockFetcher.On("FetchCountry", "India").Return(mockCountryResponse(), nil)

	var wg sync.WaitGroup
	numRequests := 50
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			country, err := service.GetCountry("India")
			assert.NoError(t, err)
			assert.Equal(t, "India", country.Name)
		}()
	}

	wg.Wait()
}
