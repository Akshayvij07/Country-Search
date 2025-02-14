package di

import (
	"github.com/Akshayvij07/country-search/internals/api"
	"github.com/Akshayvij07/country-search/internals/api/handler"
	"github.com/Akshayvij07/country-search/internals/cache"
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/Akshayvij07/country-search/internals/services"
	thirdparty "github.com/Akshayvij07/country-search/internals/third_party"
)

// ConfigureServer configures the server
// Dependency Injection is done here
func ConfigureServer() *api.Server {
	cache := setupCache()
	client := &thirdparty.APIClient{}
	service := services.New(cache, client)
	handler := handler.New(service)

	return api.NewServer("8000", handler)
}

// setupCache sets up the cache.Add a country to the cache
func setupCache() *cache.MapCache {
	cache := cache.NewMapCache()
	countryData := models.NewCountry("India", "New Delhi", "INR", 10938800)
	// Store in cache
	cache.Set("india", countryData)

	return cache
}
