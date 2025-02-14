package services

import (
	"github.com/Akshayvij07/country-search/internals/cache"
	"github.com/Akshayvij07/country-search/internals/models"
	thirdparty "github.com/Akshayvij07/country-search/internals/third_party"
	"github.com/Akshayvij07/country-search/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Services provides a layer of abstraction between the API and the cache and third-party API layers.

type Services interface {
	// GetCountry returns a country object from the cache or third-party API.
	GetCountry(name string) (models.Country, error)
}

// Service implements Services.
type Service struct {
	cache        cache.Cache
	fetchCountry thirdparty.CountryFetcher
}

// New initializes a Service object.
func New(cache cache.Cache, fetch thirdparty.CountryFetcher) *Service {
	return &Service{
		cache:        cache,
		fetchCountry: fetch,
	}
}

// GetCountry returns a country object from the cache or third-party API.
// If the country is not found in the cache, it fetches it from the third-party API.
func (s *Service) GetCountry(name string) (models.Country, error) {
	err := ValidateCountryName(name)
	if err != nil {
		log.Error().Err(err).Msg("Invalid country name")
		return models.Country{}, err
	}

	log.Info().Msg("Preparing country data...")
	c, err := s.cache.Get(name)
	if err != nil {
		if err == errors.ErrKeyNotFound {
			log.Info().Msg("Country not found in cache, fetching from third-party API...")

			country, err := s.GetCountryFromThirdParty(name)
			if err != nil {
				return models.Country{}, err
			}

			s.cache.Set(name, country)
			return country, nil
		}
		return models.Country{}, err
	}

	log.Trace().Msgf("country: %v found in cache", c)
	country := c.(models.Country)
	return country, nil
}

// GetCountryFromThirdParty fetches a country object from the third-party API.
func (s *Service) GetCountryFromThirdParty(name string) (models.Country, error) {
	data, err := thirdparty.FetchCountry(name)
	if err != nil {
		return models.Country{}, err
	}

	log.Info().Msgf("country: %v fetched from third-party API", data)

	var currency string
	for key := range data.Currencies {
		currency = key
	}

	return models.NewCountry(data.Name.Common, data.Capital[0], currency, data.Population), nil
}
