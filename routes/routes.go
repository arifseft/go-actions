package routes

import (
	"github.com/arifseft/go-actions/controllers"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	controller := controllers.UserController{}
	g.GET("/name", controller.GetName)
	g.GET("/biodata", controller.GetBiodata)
}
