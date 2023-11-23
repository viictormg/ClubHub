package usecase

import ports "github.com/viictormg/clubHub/internal/application/port/franchise"

type FranchiseUsecase struct {
	FranchiseAdapterHTTP ports.IFranchiseAdapterHTTP
	FranchiseAdapterDB   ports.IFranchiseAdapterDB
}

func NewFranchiseUsecase(FranchiseAdapter ports.IFranchiseAdapterHTTP, FranchiseAdapterDB ports.IFranchiseAdapterDB) *FranchiseUsecase {
	return &FranchiseUsecase{
		FranchiseAdapterHTTP: FranchiseAdapter,
		FranchiseAdapterDB:   FranchiseAdapterDB,
	}
}
