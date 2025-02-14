package di

import (
	"github.com/Akshayvij07/country-search/internals/api"
	"github.com/Akshayvij07/country-search/internals/api/handler"
	"github.com/Akshayvij07/country-search/internals/cache"
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/Akshayvij07/country-search/internals/services"
	thirdparty "github.com/Akshayvij07/country-search/internals/third_party"
)

func ConfigureServer() *api.Server {
	cache := setupCache()
	client := &thirdparty.APIClient{}
	service := services.New(cache, client)
	handler := handler.New(service)

	return api.NewServer("8000", handler)
}

func setupCache() *cache.MapCache {
	cache := cache.NewMapCache()
	countryData := models.NewCountry("India", "New Delhi", "INR", 10938800)
	// Store in cache
	cache.Set("india", countryData)

	return cache
}
