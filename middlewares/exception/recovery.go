package exception

import "github.com/gin-gonic/gin"

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithoutWriter(f)
}

func RecoveryWithoutWriter(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				f(c, err)
			}
		}()
		c.Next()
	}
}
