package app

import (
	"github.com/arifseft/go-actions/middlewares/exception"
	"github.com/arifseft/go-actions/routes"
	"github.com/gin-gonic/gin"
)

type Application struct {
}

func (a Application) CreateApp(r *gin.Engine) {
	r.Use(exception.Recovery(exception.ErrorHandler))
	configureAPIEndpoint(r)
}

func configureAPIEndpoint(r *gin.Engine) {
	g := r.Group("/user")
	routes.Router(g)
}
