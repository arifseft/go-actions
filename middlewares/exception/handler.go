package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	res := Exception{}
	mapstructure.Decode(err, &res)

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status": res.Status,
		"flag":   res.Flag,
		"errors": map[string]interface{}{
			"message": res.Errors.Message,
			"flag":    res.Errors.Flag,
		},
	})
	return
}

type Exception struct {
	Status int         `json:"status"`
	Flag   string      `json:"flag"`
	Errors ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Flag    string `json:"flag"`
	Message string `json:"message"`
}
