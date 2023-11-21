package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	franchiseUsecases "github.com/viictormg/clubHub/internal/application/usecase/franchise"
	franchiseAdapters "github.com/viictormg/clubHub/internal/infrastructure/adapters/http-consumer/franchise"
	franchiseHandlers "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/api/franchise"
)

func main() {
	franchiseAdapter := franchiseAdapters.NewFranchiseAdapter(&http.Client{})
	franchiseUsecase := franchiseUsecases.NewFranchiseUsecase(franchiseAdapter)
	franchiseHandler := franchiseHandlers.NewFranchiseHandler(franchiseUsecase)

	runServer(franchiseHandler)

}

func runServer(franchiseHandler *franchiseHandlers.Franchise) {
	e := echo.New()

	api := e.Group("/api/clubhub")

	api.POST("/franchise", franchiseHandler.CreateFranchiseHandler)

	err := e.Start(":3000")

	if err != nil {
		fmt.Println(err)
	}
}
