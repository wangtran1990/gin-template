package middleware

import (
	helper "template/Helper"

	"github.com/gin-gonic/gin"
)

// Middleware1 ...
func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		_procMiddleware1(c)
		c.Next()
	}
}

func _procMiddleware1(c *gin.Context) {
	helper.Logger("", "").Infoln("Middleware 1")
}
