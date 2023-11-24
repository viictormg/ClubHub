package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	pkg "github.com/viictormg/clubHub/internal/infrastructure/pkg"

	companyUsecases "github.com/viictormg/clubHub/internal/application/usecase/company"
	franchiseUsecases "github.com/viictormg/clubHub/internal/application/usecase/franchise"
	companyAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/database/company"
	countryAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/database/country"
	franchiseAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/database/franchise"
	ownerAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/database/owner"
	franchiseAdaptersHTTP "github.com/viictormg/clubHub/internal/infrastructure/adapters/http-consumer/franchise"
	companyHandlers "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/api/company"
	franchiseHandlers "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/api/franchise"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	db := pkg.NewPostgresConection()

	countryAdapter := countryAdapters.NewCountryAdapterDB(db)
	franchiseAdapter := franchiseAdapters.NewFranchiseAdapter(db)
	ownerAdapter := ownerAdapters.NewOwnerAdapter(db)
	companyAdapter := companyAdapters.NewCompanyAdapter(db)

	franchiseAdapterHTTP := franchiseAdaptersHTTP.NewFranchiseAdapterHTTP(&http.Client{})
	franchiseUsecase := franchiseUsecases.NewFranchiseUsecase(franchiseAdapterHTTP, franchiseAdapter, countryAdapter)
	franchiseHandler := franchiseHandlers.NewFranchiseHandler(franchiseUsecase)
	companyUsecase := companyUsecases.NewCompanyUsecase(countryAdapter, ownerAdapter, companyAdapter, franchiseAdapter)

	companyHandler := companyHandlers.NewCompanyHandler(companyUsecase)

	runServer(franchiseHandler, companyHandler)

}

func runServer(franchiseHandler *franchiseHandlers.Franchise, companyHandler *companyHandlers.CompanyHandler) {
	e := echo.New()

	api := e.Group("/api/clubhub")

	franchiseGroup := api.Group("/franchise")

	franchiseGroup.POST("/create", franchiseHandler.CreateFranchiseHandler)
	franchiseGroup.GET("/getByID/:id", franchiseHandler.GetFranchiseByIDHandler)
	franchiseGroup.GET("/getByParam", franchiseHandler.GetFranchiseByParamHandler)
	franchiseGroup.GET("/getByName", franchiseHandler.GetFranchiseByNameHandler)

	companyGroup := api.Group("/company")
	companyGroup.GET("/getByName", companyHandler.GetCompanyByNameHandler)

	err := e.Start(":3000")

	if err != nil {
		fmt.Println(err)
	}
}
