package routes

import (
	"github.com/arifseft/go-actions/controllers"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	g.GET("/name", controllers.GetName)
	g.GET("/biodata", controllers.GetBiodata)
}
