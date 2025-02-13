package handler

import (
	"net/http"

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

func NewHandler(service services.Services) *Handler {
	return &Handler{service: service}
}

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

	country, err := h.service.GetCountry(query.Name)
	handleResponse(c, country, err)
}

func handleResponse(c *gin.Context, data any, err error) {
	resp := buildResponse(data, err)
	c.JSON(resp.Status, resp)
}

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

var errorResponse = map[error]response.Response{
	errors.ErrKeyNotFound: response.New(http.StatusNotFound, "country not found", nil, errors.ErrKeyNotFound.Error()),
}
