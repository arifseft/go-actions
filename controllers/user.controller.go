package controllers

import (
	"net/http"

	"github.com/arifseft/go-actions/database/entity"
	"github.com/arifseft/go-actions/global/types"
	"github.com/arifseft/go-actions/middlewares/exception"
	"github.com/arifseft/go-actions/validation"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) GetName(c *gin.Context) {
	name := "M Arif Sefrianto"
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get name",
		"result":  name,
	})
}

func (u UserController) GetBiodata(c *gin.Context) {
	biodata := types.GetBiodataResult{
		Name:    "M Arif Sefrianto",
		Address: "Lubuklinggau",
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get biodata",
		"result":  biodata,
	})
}

func (u UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		exception.BadRequest(err.Error(), "ERROR_BIND_REQUEST_JSON")
	}

	userValidate := &validation.CreateUserSchema{Name: user.Name}
	validation.Validate(userValidate)

	data := types.CreateUserResult{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Email:   user.Email,
		Address: user.Address,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
		"result":  data,
	})
}
