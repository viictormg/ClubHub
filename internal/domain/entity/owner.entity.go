package entity

type OwnerEntity struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	CountryCodeID int    `json:"countryCodeId"`
	Address       string `json:"address"`
	ZipCode       string `json:"zipCode"`
}
