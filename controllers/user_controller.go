package controllers

import (
	"fmt"
	"net/http"

	"github.com/arifseft/go-actions/database/entity"
	"github.com/arifseft/go-actions/global/types"
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
	user := new(entity.User)
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
	}

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
