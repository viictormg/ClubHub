package usecase

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/clubHub/internal/application/mappers"
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

const (
	errorGetingCompany string = "error getting company"
)

func (c *CompanyUsecase) GetCompanyByNameUsecase(name string) (*dto.ResponseCompanyByNameDTO, error) {
	var companyResponse dto.ResponseCompanyByNameDTO
	var franchisesDTP []dto.FranchiseDTO
	company, err := c.companyAdapter.GetCompanyByNameAdapter(name)
	if company == nil || err != nil {
		logrus.Error(err)
		return nil, errors.New(errorGetingCompany)
	}

	owner, err := c.ownerAdapter.GetOwnerByCompanyIDAdapter(company.ID)
	if owner == nil || err != nil {
		logrus.Error(err)
		return nil, errors.New(errorGetingCompany)
	}
	franchises, err := c.franchiseAdapter.GetFranchisesByCompanyID(company.ID)
	if owner == nil || err != nil {
		logrus.Error(err)
		return nil, errors.New(errorGetingCompany)
	}

	fmt.Println(franchises)

	mappers.MapCompanyEntityToResposeCompanyDTO(company, &companyResponse)
	mappers.MapOwnerEntityToResponseCompany(owner, &companyResponse)

	for _, franchiseEntity := range *franchises {
		newFranchiseDTO := mappers.MapFranchiseEntityToFranchiseDTO(&franchiseEntity, &entity.CountryEntity{})
		franchisesDTP = append(franchisesDTP, *newFranchiseDTO)
	}
	companyResponse.Company.Franchises = franchisesDTP

	fmt.Println(owner)
	return &companyResponse, nil
}
