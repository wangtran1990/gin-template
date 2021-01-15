package middleware

import (
	"fmt"
	helper "template/Helper"

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
	txnID := helper.LogTraceID()
	c.Set("x-request-id", txnID)                 // set to request param
	c.Writer.Header().Set("X-Request-Id", txnID) // set to request header

	curl, _ := http2curl.GetCurlCommand(c.Request)
	helper.Logger(txnID, "").Infoln(fmt.Sprint(curl))
}
