package middleware

import (
	services "template/Services"

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
	services.Logger("", "").Infoln("Middleware 2")
}
