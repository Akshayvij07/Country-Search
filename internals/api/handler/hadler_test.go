package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Akshayvij07/country-search/internals/api/handler"
	"github.com/Akshayvij07/country-search/internals/mocks"
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test_GetCountry tests the GetCountry handler function
func Test_GetCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockService)
	handler := handler.New(mockService)

	expectedCountry := models.NewCountry("India", "New Delhi", "INR", 1380004385)

	// Mock service response
	mockService.On("GetCountry", "India").Return(expectedCountry, nil)

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/country?name=India", nil)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	handler.GetCountry(c)

	// Assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
	mockService.AssertExpectations(t)

}

// Test_InvalidCountryRequest tests the GetCountry handler function
func Test_InvalidCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockService)
	handler := handler.New(mockService)

	// Mock service response
	mockService.On("GetCountry", "Unknown").Return(models.Country{}, assert.AnError)

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/country?name=Unknown", nil)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	handler.GetCountry(c)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	mockService.AssertExpectations(t)

}
