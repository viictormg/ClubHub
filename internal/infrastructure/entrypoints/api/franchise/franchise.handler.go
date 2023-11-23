package api

import "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/ports"

type Franchise struct {
	franchiseUsecase ports.IFranchiseUsecase
}

func NewFranchiseHandler(franchiseUsecase ports.IFranchiseUsecase) *Franchise {
	return &Franchise{
		franchiseUsecase: franchiseUsecase,
	}
}
