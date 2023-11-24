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
	GetFranchiseByIDAdapter(id int) (*entity.FranchiseEntity, error)
	GetFranchiseByNameAdapter(name string) (*entity.FranchiseEntity, error)
	GetFranchisesByParamAdapter(key, value string) (*[]entity.FranchiseEntity, error)
	GetFranchisesByCompanyID(companyID int) (*[]entity.FranchiseEntity, error)
}
