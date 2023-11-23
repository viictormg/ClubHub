package entity

type RedirectTo struct {
	IPAddress  string `json:"ipAddress"`
	ServerName string `json:"serverName"`
}

type FranchiseEntity struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	City               string `json:"city"`
	CountryCodeID      int    `json:"countryCodeID"`
	Address            string `json:"address"`
	ZipCode            string `json:"zipCode"`
	CompanyID          int    `json:"companyID"`
	Protocol           string `json:"protocol"`
	LogoURL            string `json:"logoURL"`
	DomainCreation     string `json:"domainCreation"`
	DomainExpiration   string `json:"domainExpiration"`
	DomainOwnerName    string `json:"domainOwnerName"`
	DomainContactEmail string `json:"domainContactEmail"`
	Redirections       string `json:"redirections"`
	NumberRedirections int    `json:"numberRedirections"`
}
