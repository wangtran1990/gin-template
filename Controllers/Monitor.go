//Controllers/Monitor.go

package controllers

import (
	"fmt"
	"net/http"
	helper "template/Helper"

	"github.com/gin-gonic/gin"
)

// HealthCheck ... health check service
func HealthCheck(c *gin.Context) {
	resultDB := helper.CheckDatabase()
	if resultDB == "" {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	c.JSON(http.StatusOK, "DB time = "+resultDB)
}

// LoadTest ...
func LoadTest(c *gin.Context) {

	// #1
	for i := 0; i < 10; i++ {
		tmp := i
		fmt.Println(tmp)
	}

	// #2
	resultDB := helper.CheckDatabase()
	fmt.Println(resultDB)

	// #3
	tmp := 0
	for i := 0; i < 10000; i++ {
		tmp += i
	}
	fmt.Println(tmp)

	// #4
	resultDB = helper.CheckDatabase()
	fmt.Println(resultDB)

	c.JSON(http.StatusOK, "OK")
}
