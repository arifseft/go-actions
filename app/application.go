package app

import (
	"github.com/arifseft/go-actions/routes"
	"github.com/labstack/echo/v4"
)

type Application struct {
}

func (a Application) CreateApp(e *echo.Echo) {
	configureAPIEndpoint(e)
}

func configureAPIEndpoint(e *echo.Echo) {
	g := e.Group("/user")
	routes.Router(g)

}
