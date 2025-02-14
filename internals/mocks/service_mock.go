package mocks

import (
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetCountry(name string) (models.Country, error) {
	args := m.Called(name)
	return args.Get(0).(models.Country), args.Error(1)
}
