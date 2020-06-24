package routes

import (
	"github.com/arifseft/go-actions/controllers"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	controller := controllers.UserController{}
	{
		g.GET("/name", controller.GetName)
		g.GET("/biodata", controller.GetBiodata)
		g.POST("/user", controller.CreateUser)
	}
}
