package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	pkg "github.com/viictormg/clubHub/internal/infrastructure/pkg"

	franchiseUsecases "github.com/viictormg/clubHub/internal/application/usecase/franchise"
	franchiseAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/database/franchise"
	franchiseAdaptersHTTP "github.com/viictormg/clubHub/internal/infrastructure/adapters/http-consumer/franchise"
	franchiseHandlers "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/api/franchise"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	db := pkg.NewPostgresConection()

	franchiseAdapter := franchiseAdapters.NewFranchiseAdapter(db)

	franchiseAdapterHTTP := franchiseAdaptersHTTP.NewFranchiseAdapterHTTP(&http.Client{})
	franchiseUsecase := franchiseUsecases.NewFranchiseUsecase(franchiseAdapterHTTP, franchiseAdapter)
	franchiseHandler := franchiseHandlers.NewFranchiseHandler(franchiseUsecase)

	runServer(franchiseHandler)

}

func runServer(franchiseHandler *franchiseHandlers.Franchise) {
	e := echo.New()

	api := e.Group("/api/clubhub")

	franchiseGroup := api.Group("/franchise")

	franchiseGroup.POST("/create", franchiseHandler.CreateFranchiseHandler)
	franchiseGroup.GET("/getByID/:id", franchiseHandler.GetFranchiseByIDHandler)

	err := e.Start(":3000")

	if err != nil {
		fmt.Println(err)
	}
}
