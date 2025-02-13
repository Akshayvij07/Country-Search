package services

import (
	"github.com/Akshayvij07/country-search/internals/cache"
	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/rs/zerolog/log"
)

type Services interface {
	GetCountry(name string) (models.Country, error)
}

type Service struct {
	cache cache.Cache
}

func NewService(cache cache.Cache) *Service {
	return &Service{cache: cache}
}

func (s *Service) GetCountry(name string) (models.Country, error) {
	data, err := s.cache.Get(name)
	if err != nil {
		return models.Country{}, err
	}

	log.Info().Msgf("country: %v", data)
	country := data.(models.Country)
	return country, nil
}
