package routes

import (
	"github.com/Akshayvij07/country-search/internals/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup, handler *handler.Handler) {
	router.GET("/search", handler.GetCountry)
}
