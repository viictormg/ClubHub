package entity

type Company struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	TaxNumber     string `json:"taxNumber"`
	City          string `json:"city"`
	CountryCodeID int    `json:"countryCodeId"`
	OwnerID       int    `json:"ownerId"`
	Address       string `json:"address"`
	ZipCode       string `json:"zipCode"`
}
