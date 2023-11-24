package api

import "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/ports"

type CompanyHandler struct {
	companyUsecase ports.ICompanyUsecase
}

func NewCompanyHandler(companyUsecase ports.ICompanyUsecase) *CompanyHandler {
	return &CompanyHandler{
		companyUsecase: companyUsecase,
	}
}
