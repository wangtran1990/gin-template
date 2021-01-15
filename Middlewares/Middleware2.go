package middleware

import (
	helper "template/Helper"

	"github.com/gin-gonic/gin"
)

// Middleware2 ...
func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		_procMiddleware2(c)
		c.Next()
	}
}

func _procMiddleware2(c *gin.Context) {
	helper.Logger("", "").Infoln("Middleware 2")
}
