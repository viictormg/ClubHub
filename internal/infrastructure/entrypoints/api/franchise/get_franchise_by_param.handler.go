package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	messageParamaRequired string = "%s is required"
)

func (f *Franchise) GetFranchiseByParamHandler(c echo.Context) error {

	key := c.QueryParam("key")
	if key == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf(messageParamaRequired, "key"))
	}
	value := c.QueryParam("value")

	if value == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf(messageParamaRequired, "value"))
	}

	response, err := f.franchiseUsecase.GetFranchiseParamUsecase(key, value)
	if err != nil {
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, response)

}
