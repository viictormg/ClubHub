package ports

import "github.com/viictormg/clubHub/internal/domain/dto"

type IFranchiseAdapter interface {
	ExtractURLInfoAdapter(url string) (*dto.SSLInfoResultDTO, error)
	ExtractInfoDomainAdapter(url string) (*string, error)
	ExtractAssetsPageAdapter(url string) (*string, error)
}
