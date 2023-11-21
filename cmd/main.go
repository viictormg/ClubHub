package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	franchiseHandlers "github.com/viictormg/clubHub/internal/infrastructure/entrypoints/api/franchise"
)

func main() {

	franchiseHandler := franchiseHandlers.NewFranchise()

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
