package ports

import "github.com/viictormg/clubHub/internal/domain/entity"

type ICompanyAdapter interface {
	GetCompanyByNameAdapter(name string) (*entity.CompanyEntity, error)
}
