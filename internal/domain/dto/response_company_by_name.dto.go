package dto

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Address string `json:"address"`
	ZipCode string `json:"zipCode"`
}

type Contact struct {
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Location Location `json:"location"`
}

type Owner struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Contact   Contact `json:"contact"`
}

type Information struct {
	Name      string   `json:"name"`
	TaxNumber string   `json:"tax_number"`
	Location  Location `json:"location"`
}

type CompanyData struct {
	Owner       Owner          `json:"owner"`
	Information Information    `json:"informacion"`
	Franchises  []FranchiseDTO `json:"franchises"`
}

type ResponseCompanyByNameDTO struct {
	Company CompanyData `json:"company"`
}
