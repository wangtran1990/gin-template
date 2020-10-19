//Controllers/User.go

package controllers

import (
	"fmt"
	"net/http"
	models "template/models"
	services "template/services"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	requestID := c.GetString("x-request-id")
	services.Logger(requestID, "").Infoln("RequestID= ", requestID)
	// cacheTest := services.CacheExists("xxxxxxxxxx")
	// services.Logger(requestID, "").Infoln("cacheTest= ", cacheTest)

	httpCode, body, erro := services.MakeHTTPRequest("GET", "https://api-101.glitch.me/customers", "", nil, true)
	services.Logger(requestID, "").Infoln("httpCode= ", httpCode)
	services.Logger(requestID, "").Infoln("body= ", fmt.Sprintf("%s", body))
	services.Logger(requestID, "").Infoln("error= ", erro)

	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
