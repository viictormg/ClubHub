package usecase

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/clubHub/internal/application/mappers"
	"github.com/viictormg/clubHub/internal/domain/dto"
)

const (
	errGettingFranchise = "error getting franchise"
)

func (f *FranchiseUsecase) GetFranchiseByNameUsecase(name string) (*dto.FranchiseDTO, error) {
	response, err := f.FranchiseAdapterDB.GetFranchiseByNameAdapter(name)

	if err != nil || response == nil {
		logrus.Error(err)
		return nil, errors.New(errGettingFranchise)
	}
	country, err := f.CountryCodeAdapter.GetCountryByIDAdapter(response.CountryCodeID)
	if err != nil || response == nil {
		logrus.Error(err)
		return nil, errors.New(errGettingFranchise)
	}
	return mappers.MapFranchiseEntityToFranchiseDTO(response, country), nil
}
