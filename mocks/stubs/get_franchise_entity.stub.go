package stubs

import "github.com/viictormg/clubHub/internal/domain/entity"

func GetFranchiseEntity() entity.FranchiseEntity {
	return entity.FranchiseEntity{
		ID:                 0,
		URL:                "marriot.com",
		City:               "MEDELLIN",
		CountryCodeID:      1,
		LogoURL:            "http://favico.png",
		Redirections:       "[{\"ipAddress\":\"1\",\"serverName\":\"\"}]",
		NumberRedirections: 1,
		CompanyID:          1,
	}
}
