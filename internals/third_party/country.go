package thirdparty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Akshayvij07/country-search/internals/models"
	"github.com/Akshayvij07/country-search/pkg/errors"
)

type CountryFetcher interface {
	FetchCountry(name string) (*models.CountryResponse, error)
}

// Implement the interface in the actual struct
type APIClient struct{}

func (c *APIClient) FetchCountry(name string) (*models.CountryResponse, error) {
	return FetchCountry(name) // Call the actual function
}

func FetchCountry(name string) (*models.CountryResponse, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", name)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.ErrKeyNotFound
		}
		return nil, fmt.Errorf("API returned non-OK status: %d", resp.StatusCode)
	}

	// Parse JSON response
	var countries []models.CountryResponse
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("no country found for %s", name)
	}

	if err := validateResponse(&countries[0]); err != nil {
		return nil, err
	}

	return &countries[0], nil
}

func validateResponse(response *models.CountryResponse) error {
	if response.Name.Common == "" {
		return fmt.Errorf("country name not found in response")
	}
	if response.Capital == nil || len(response.Capital) == 0 {
		return fmt.Errorf("country capital not found in response")
	}
	if response.Currencies == nil || len(response.Currencies) == 0 {
		return fmt.Errorf("country currency not found in response")
	}
	if response.Population == 0 {
		return fmt.Errorf("country population not found in response")
	}
	return nil
}
