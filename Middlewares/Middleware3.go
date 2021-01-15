package middleware

import (
	helper "template/Helper"

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
	helper.Logger("", "").Infoln("Middleware 3")
}
