package main

import (
	"github.com/arifseft/go-actions/app"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	app := new(app.Application)
	app.CreateApp(e)

	e.Logger.Fatal(e.Start(":8080"))
}
