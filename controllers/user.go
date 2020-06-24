package controllers

import (
	"net/http"

	"github.com/arifseft/go-actions/utils"
	"github.com/labstack/echo/v4"
)

type GetNameResponse struct {
	utils.Response
	Result string `json:"result"`
}

type GetBiodataResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetBiodataResponse struct {
	utils.Response
	Result GetBiodataResult `json:"result"`
}

func GetName(c echo.Context) error {
	name := GetNameResponse{
		Result: "M Arif Sefrianto",
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success get name",
		},
	}

	return c.JSON(http.StatusOK, name)
}

func GetBiodata(c echo.Context) error {
	biodata := GetBiodataResponse{
		Response: utils.Response{
			Status: http.StatusOK,
			Message: "Success to get biodata",
		},
		Result: GetBiodataResult{
			Name: "M Arif Sefrianto",
			Address: "Lubuklinggau",
		},
	}

	return c.JSON(http.StatusOK, biodata)
}
