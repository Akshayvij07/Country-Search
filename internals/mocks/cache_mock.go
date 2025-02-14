package mocks

import "github.com/stretchr/testify/mock"

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0), args.Error(1)
}

func (m *MockCache) Set(key string, value interface{}) {
	m.Called(key, value)
}
