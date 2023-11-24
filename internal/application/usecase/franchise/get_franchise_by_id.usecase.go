package usecase

import (
	"github.com/viictormg/clubHub/internal/application/mappers"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

func (f *FranchiseUsecase) GetFranchiseByIDUsecase(id int) (*dto.FranchiseDTO, error) {
	response, err := f.FranchiseAdapterDB.GetFranchiseByIDAdapter(id)
	if err != nil {
		return nil, err
	}

	country, err := f.CountryCodeAdapter.GetCountryByIDAdapter(response.CountryCodeID)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return mappers.MapFranchiseEntityToFranchiseDTO(response, country), nil
}
