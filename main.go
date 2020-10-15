//main.go
package main

import (
	"os"
	"runtime"
	configs "template/Configs"
	routes "template/Routes"
	services "template/Services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
)

var err error
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	// Set run CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	services.Logger("", "").Infoln("Num CPU: ", runtime.NumCPU())

	// Load the .env file in the current directory
	// .env already load in-case run by docker-compose
	err = godotenv.Load()
	if err != nil {
		services.Logger("", "").Warningln("Cannot get config from .env file manual - if run from docker-compose skip this warning", err)
	} else {
		services.Logger("", "").Infoln("Get config from .env successful")
	}

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
}

func main() {
	// Init routes
	r := routes.SetupRouter()

	// Start server
	runport := os.Getenv("RUN_PORT")
	if runport == "" {
		runport = "2000"
	}
	r.Run(":" + runport) // running
}
