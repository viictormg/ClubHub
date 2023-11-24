package usecase

import ports "github.com/viictormg/clubHub/internal/application/port"

type CompanyUsecase struct {
	countryAdapter   ports.ICountryCodeAdapterDB
	ownerAdapter     ports.IOwnerAdapter
	companyAdapter   ports.ICompanyAdapter
	franchiseAdapter ports.IFranchiseAdapterDB
}

func NewCompanyUsecase(
	countryAdapter ports.ICountryCodeAdapterDB,
	ownerAdapter ports.IOwnerAdapter,
	companyAdapter ports.ICompanyAdapter,
	franchiseAdapter ports.IFranchiseAdapterDB,
) *CompanyUsecase {
	return &CompanyUsecase{
		countryAdapter:   countryAdapter,
		ownerAdapter:     ownerAdapter,
		companyAdapter:   companyAdapter,
		franchiseAdapter: franchiseAdapter,
	}
}
