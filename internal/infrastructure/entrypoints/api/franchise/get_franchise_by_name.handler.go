package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (f *Franchise) GetFranchiseByNameHandler(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf(messageParamaRequired, "name"))
	}

	response, err := f.franchiseUsecase.GetFranchiseByNameUsecase(name)
	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, response)

}
