package controllers

import (
	"net/http"

	"github.com/arifseft/go-actions/global/types"
	"github.com/arifseft/go-actions/utils"
	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func (u UserController) GetName(c echo.Context) error {
	name := types.GetNameResponse{
		Result: "M Arif Sefrianto",
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success get name",
		},
	}

	return c.JSON(http.StatusOK, name)
}

func (u UserController) GetBiodata(c echo.Context) error {
	biodata := types.GetBiodataResponse{
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success to get biodata",
		},
		Result: types.GetBiodataResult{
			Name:    "M Arif Sefrianto",
			Address: "Lubuklinggau",
		},
	}

	return c.JSON(http.StatusOK, biodata)
}
