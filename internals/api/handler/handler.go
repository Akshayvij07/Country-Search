package handler

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/Akshayvij07/country-search/pkg/errors"

	"github.com/Akshayvij07/country-search/internals/api/request"
	"github.com/Akshayvij07/country-search/internals/api/response"
	"github.com/Akshayvij07/country-search/internals/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.Services
}

func New(service services.Services) *Handler {
	return &Handler{service: service}
}

// Handler function for fetching countryDetails
func (h *Handler) GetCountry(c *gin.Context) {
	var query request.CountryQuery

	if len(c.Request.URL.Query()) > 0 {
		err := c.ShouldBindQuery(&query)
		if err != nil {
			errResponse := response.BindQueryErr(err)
			c.JSON(errResponse.Status, errResponse)
			return
		}
	}

	log.Info().Msgf("query: %v", query)

	country, err := h.service.GetCountry(strings.ToLower(query.Name))
	handleResponse(c, country, err)
}

// Helper function to handle response
func handleResponse(c *gin.Context, data any, err error) {
	resp := buildResponse(data, err)
	c.JSON(resp.Status, resp)
}

// Helper function to build response
func buildResponse(data any, err error) *response.Response {
	var resp *response.Response
	if err != nil {
		if response, ok := errorResponse[err]; ok {
			resp = &response
			return resp
		}
		resp = &response.Response{
			Status:  http.StatusInternalServerError,
			Message: "unexpected error caused by server",
			Error:   err.Error(),
		}
	} else {
		resp = &response.Response{
			Status:  http.StatusOK,
			Data:    data,
			Message: "success",
		}
	}
	return resp
}

// Map of error to response.The errorResponse map is used to map errors to their corresponding response
var errorResponse = map[error]response.Response{
	errors.ErrKeyNotFound: response.New(http.StatusNotFound, "country not found", nil, errors.ErrKeyNotFound.Error()),
}
