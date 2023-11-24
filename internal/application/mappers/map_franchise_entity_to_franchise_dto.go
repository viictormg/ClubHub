package mappers

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func MapFranchiseEntityToFranchiseDTO(franchise *entity.FranchiseEntity, countryCode *entity.CountryEntity) *dto.FranchiseDTO {
	return &dto.FranchiseDTO{
		Name: franchise.Name,
		URL:  franchise.URL,
		Location: dto.LocationFranchiseDTO{
			City:          franchise.City,
			CountryCodeID: countryCode.Name,
		},
	}
}
