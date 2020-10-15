//main.go
package main

import (
	"os"
	configs "template/Configs"
	routes "template/Routes"
	services "template/Services"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var err error
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {

	// Set run mode
	runmode := os.Getenv("PROD_MODE")
	if runmode == "1" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init Log
	configs.InitLogger()

	// Start DB
	err = configs.InitDB()
	if err != nil {
		services.Logger("", "").Warningln("Can't connect RDB. ", err)
	} else {
		services.Logger("", "").Infoln("Database connect successful")
		defer configs.DB.Close()

		// Auto migration
		migration := os.Getenv("RDB_AUTO_MIGRATION")
		if migration == "1" {
			services.MigrateDataTable()
		}
	}

	// Start Cache
	err = configs.InitCache()
	if err != nil {
		services.Logger("", "").Warningln("Can't connect Cache service. ", err)
	} else {
		services.Logger("", "").Infoln("Cache server connect successful")
	}

	r := routes.SetupRouter()
	runport := os.Getenv("RUN_PORT")
	if runport == "" {
		runport = "2000"
	}
	r.Run(":" + runport) // running
}
