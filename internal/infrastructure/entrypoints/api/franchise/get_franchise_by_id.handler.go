package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (f *Franchise) GetFranchiseByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	response, err := f.franchiseUsecase.GetFranchiseByIDUsecase(id)
	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, response)

}
