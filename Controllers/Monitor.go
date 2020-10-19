//Controllers/Monitor.go

package controllers

import (
	"net/http"
	services "template/services"

	"github.com/gin-gonic/gin"
)

// HealthCheck ... health check service
func HealthCheck(c *gin.Context) {
	resultDB := services.CheckDatabase()
	if resultDB == "" {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	c.JSON(http.StatusOK, "DB time = "+resultDB)
}
