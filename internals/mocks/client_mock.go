package mocks

import (
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/stretchr/testify/mock"
)

// Mock thirdparty
type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) FetchCountry(name string) (*models.CountryResponse, error) {
	args := m.Called(name)
	return args.Get(0).(*models.CountryResponse), args.Error(1)
}
