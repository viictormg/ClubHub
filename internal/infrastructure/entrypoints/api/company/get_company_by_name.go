package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/viictormg/clubHub/internal/domain/constants"
)

func (h *CompanyHandler) GetCompanyByNameHandler(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf(constants.MessageParamaRequired, "name"))
	}
	response, err := h.companyUsecase.GetCompanyByNameUsecase(name)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusConflict, err)
	}
	return c.JSON(http.StatusOK, response)
}
