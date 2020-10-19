package middleware

import (
	"fmt"
	services "template/services"

	"github.com/gin-gonic/gin"
	"moul.io/http2curl"
)

// RequestLog ...
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		_procRequestLog(c)
		c.Next()
	}
}

func _procRequestLog(c *gin.Context) {
	txnID := services.LogTraceID()
	c.Set("x-request-id", txnID)                 // set to request param
	c.Writer.Header().Set("X-Request-Id", txnID) // set to request header

	curl, _ := http2curl.GetCurlCommand(c.Request)
	services.Logger(txnID, "").Infoln(fmt.Sprint(curl))
}
