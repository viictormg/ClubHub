package dto

type LocationFranchiseDTO struct {
	City          string `json:"city"`
	CountryCodeID string `json:"countryCode"`
	Address       string `json:"address"`
	ZipCode       string `json:"zipCode"`
}
type FranchiseDTO struct {
	Name     string               `json:"name"`
	URL      string               `json:"url"`
	Location LocationFranchiseDTO `json:"location"`
}
