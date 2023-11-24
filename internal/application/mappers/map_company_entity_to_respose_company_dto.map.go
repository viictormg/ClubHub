package mappers

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func MapCompanyEntityToResposeCompanyDTO(companyEntity *entity.CompanyEntity, response *dto.ResponseCompanyByNameDTO) {
	response.Company.Information.Name = companyEntity.Name
	response.Company.Information.TaxNumber = companyEntity.TaxNumber
	response.Company.Information.Location.City = companyEntity.City
	response.Company.Information.Location.Country = companyEntity.City
	response.Company.Information.Location.Address = companyEntity.Address
	response.Company.Information.Location.ZipCode = companyEntity.ZipCode
}
