package usecase

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/clubHub/internal/application/mappers"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

func (f *FranchiseUsecase) GetFranchiseParamUsecase(key, value string) (*[]dto.FranchiseDTO, error) {
	var franchises []dto.FranchiseDTO
	response, err := f.FranchiseAdapterDB.GetFranchisesByParamAdapter(key, value)
	if err != nil || response == nil {
		logrus.Error(err)
		return nil, errors.New(errGettingFranchise)
	}
	countryCodeID := *response

	country, err := f.CountryCodeAdapter.GetCountryByIDAdapter(countryCodeID[0].CountryCodeID)
	if err != nil || response == nil {
		logrus.Error(err)
		return nil, errors.New(errGettingFranchise)
	}

	for _, franchiseEntity := range *response {
		newFranchiseDTO := mappers.MapFranchiseEntityToFranchiseDTO(&franchiseEntity, country)
		franchises = append(franchises, *newFranchiseDTO)
	}

	return &franchises, nil
}
