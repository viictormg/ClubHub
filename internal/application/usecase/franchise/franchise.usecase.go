package usecase

import ports "github.com/viictormg/clubHub/internal/application/port"

type FranchiseUsecase struct {
	FranchiseAdapterHTTP ports.IFranchiseAdapterHTTP
	FranchiseAdapterDB   ports.IFranchiseAdapterDB
	CountryCodeAdapter   ports.ICountryCodeAdapterDB
}

func NewFranchiseUsecase(
	FranchiseAdapter ports.IFranchiseAdapterHTTP,
	FranchiseAdapterDB ports.IFranchiseAdapterDB,
	CountryCodeAdapter ports.ICountryCodeAdapterDB,
) *FranchiseUsecase {
	return &FranchiseUsecase{
		FranchiseAdapterHTTP: FranchiseAdapter,
		FranchiseAdapterDB:   FranchiseAdapterDB,
		CountryCodeAdapter:   CountryCodeAdapter,
	}
}
