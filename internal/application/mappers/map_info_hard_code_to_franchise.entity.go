package mappers

import "github.com/viictormg/clubHub/internal/domain/entity"

func MapInfoHardCodeToFranchiseEntity(franchise *entity.FranchiseEntity) {
	franchise.CompanyID = 1
	franchise.CountryCodeID = 1
	franchise.City = "MEDELLIN"
}
