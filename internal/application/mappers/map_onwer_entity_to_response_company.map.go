package mappers

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

func MapOwnerEntityToResponseCompany(owner *entity.OwnerEntity, response *dto.ResponseCompanyByNameDTO) {
	response.Company.Owner.FirstName = owner.FirstName
	response.Company.Owner.LastName = owner.LastName
	response.Company.Owner.Contact.Email = owner.Email
	response.Company.Owner.Contact.Phone = owner.Phone
	response.Company.Owner.Contact.Location.Address = owner.Address
	response.Company.Owner.Contact.Location.City = owner.City
	response.Company.Owner.Contact.Location.City = owner.City
	response.Company.Owner.Contact.Location.ZipCode = owner.ZipCode

}
