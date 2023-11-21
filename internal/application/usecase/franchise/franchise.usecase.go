package usecase

import ports "github.com/viictormg/clubHub/internal/application/port/franchise"

type FranchiseUsecase struct {
	FranchiseAdapter ports.IFranchiseAdapter
}

func NewFranchiseUsecase(FranchiseAdapter ports.IFranchiseAdapter) *FranchiseUsecase {
	return &FranchiseUsecase{
		FranchiseAdapter: FranchiseAdapter,
	}
}
