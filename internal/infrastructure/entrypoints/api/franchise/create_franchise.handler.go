package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/clubHub/internal/application/model"
)

func (f *Franchise) CreateFranchiseHandler(c echo.Context) error {
	var franchiseModelCreate model.FranchiseCreateModel

	err := c.Bind(&franchiseModelCreate)

	fmt.Println(franchiseModelCreate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("malformed body"))
	}
	err = franchiseModelCreate.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	return nil
}
