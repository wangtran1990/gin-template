//Routes/Routes.go

package routes

import (
	controllers "template/Controllers"
	middlewares "template/Middlewares"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	routes := gin.Default()

	// Use global middleware
	routes.Use(middlewares.RequestLog())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	routes.Use(gin.Recovery())

	// Routing by group ***************
	// Demo group
	grpDemo := routes.Group("/dm")
	{
		grpDemo.GET("http_request", controllers.HTTPRequestDemo)
		grpDemo.GET("http_request_retry", controllers.HTTPRetryRequestDemo)
		grpDemo.GET("http_request_async", controllers.HTTPAsyncRequestDemo)
	}
	// Monitor group
	grpMonitor := routes.Group("/m")
	{
		grpMonitor.GET("health_check", controllers.HealthCheck)
		grpMonitor.GET("load_test", controllers.LoadTest)
	}

	// User group
	grpUser := routes.Group("/user-api")
	{
		grpUser.GET("user", middlewares.Middleware3(), controllers.GetUsers)
		grpUser.GET("user/:id", controllers.GetUserByID)
		grpUser.POST("user", controllers.CreateUser)
		grpUser.PUT("user/:id", controllers.UpdateUser)
		grpUser.DELETE("user/:id", controllers.DeleteUser)
	}
	return routes
}
