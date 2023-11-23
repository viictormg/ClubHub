package ports

import "github.com/viictormg/clubHub/internal/domain/dto"

type IFranchiseAdapterHTTP interface {
	ExtractURLInfoAdapterHTTP(url string) (*dto.SSLInfoResultDTO, error)
	ExtractInfoDomainAdapterHTTP(url string) (*string, error)
	ExtractAssetsPageAdapterHTTP(url string) (*string, error)
	ExtractFooterPageAdapterHTTP(url string) (*string, error)
}

type IFranchiseAdapterDB interface {
	GetFranchiseByID(id int) (*dto.Franchise, error)
}
