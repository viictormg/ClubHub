package ports

import "github.com/viictormg/clubHub/internal/domain/dto"

type ICompanyUsecase interface {
	GetCompanyByNameUsecase(name string) (*dto.ResponseCompanyByNameDTO, error)
}
