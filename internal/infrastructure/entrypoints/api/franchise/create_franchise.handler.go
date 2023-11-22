package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/clubHub/internal/application/model"
)

func (f *Franchise) CreateFranchiseHandler(c echo.Context) error {
	var franchiseModelCreate model.FranchiseCreateModel

	err := c.Bind(&franchiseModelCreate)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("malformed body"))
	}
	err = franchiseModelCreate.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := f.franchiseUsecase.CreateFranchiseUsecase(franchiseModelCreate)

	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, response)

}
