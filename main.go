package main

import (
	"github.com/arifseft/go-actions/app"
	"github.com/arifseft/go-actions/database"
	"github.com/arifseft/go-actions/database/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app := new(app.Application)
	app.CreateApp(r)

	database.Connection()
	db := database.GetDB()
	migration.Migrate(db)

	r.Run()
}
