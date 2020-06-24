package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	res := Exception{}
	if err := mapstructure.Decode(err, &res); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusInternalServerError,
			"flag":   "INTERNAL_SERVER_ERROR",
			"errors": map[string]interface{}{
				"message": "An error occured on our server",
				"flag":    "ERROR_MAP_TO_STRUCT",
			},
		})
		return
	}
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
