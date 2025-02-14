package services

import "github.com/Akshayvij07/country-search/pkg/errors"

func ValidateCountryName(name string) error {
	if name == "" {
		return errors.ErrParams
	}
	return nil
}
