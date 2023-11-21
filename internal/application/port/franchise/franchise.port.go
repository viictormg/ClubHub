package ports

import "github.com/viictormg/clubHub/internal/domain/dto"

type IFranchiseAdapter interface {
	ExtractURLInfoAdapter(url string) (*dto.SSLInfoResultDTO, error)
}
