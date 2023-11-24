package ports

import "github.com/viictormg/clubHub/internal/domain/entity"

type ICountryCodeAdapterDB interface {
	GetCountryByIDAdapter(id int) (*entity.CountryEntity, error)
}
