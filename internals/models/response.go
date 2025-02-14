package models

// Country represents a country.
type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Currency   string `json:"currency"`
	Population int    `json:"population"`
}

// NewCountry creates a new country.
func NewCountry(name string, capital string, currency string, population int) Country {
	return Country{Name: name, Capital: capital, Currency: currency, Population: population}
}
