package ports

import (
	"github.com/viictormg/clubHub/internal/domain/dto"
	"github.com/viictormg/clubHub/internal/domain/entity"
)

type IFranchiseAdapterHTTP interface {
	ExtractURLInfoAdapterHTTP(url string) (*dto.SSLInfoResultDTO, error)
	ExtractInfoDomainAdapterHTTP(url string) (*string, error)
	ExtractAssetsPageAdapterHTTP(url string) (*string, error)
	ExtractFooterPageAdapterHTTP(url string) (*string, error)
}

type IFranchiseAdapterDB interface {
	CreateFranchiseAdapter(newFranchise *entity.FranchiseEntity) (*dto.CreationDTO, error)
	GetFranchiseByIDAdapter(id int) (*dto.Franchise, error)
	GetFranchisesByParamAdapter(key, value string) (*[]dto.Franchise, error)
}
