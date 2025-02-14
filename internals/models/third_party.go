package models

// CountryResponse represents a country response from the third-party API.
type CountryResponse struct {
	Name       Name                `json:"name"`
	Capital    []string            `json:"capital"`
	Currencies map[string]Currency `json:"currencies"`
	Population int                 `json:"population"`
}

// Name represents a country name.
type Name struct {
	Common string `json:"common"`
}

// Currency represents a currency.
type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
