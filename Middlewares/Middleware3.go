package middleware

import (
	services "template/Services"

	"github.com/gin-gonic/gin"
)

// Middleware3 ...
func Middleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		_procMiddleware1(c)
		_procMiddleware2(c)
		_procMiddleware3(c)
		c.Next()
	}
}

func _procMiddleware3(c *gin.Context) {
	services.Logger("", "").Infoln("Middleware 3")
}
