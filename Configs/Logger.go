package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// InitLogger ...
func InitLogger() {
	// More configs at https://godoc.org/github.com/sirupsen/logrus
	/* LOG level
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.Fatal("Bye.") // Calls os.Exit(1) after logging
	log.Panic("I'm bailing.") // Calls panic() after logging
	*/

	runmode := os.Getenv("PROD_MODE")
	if runmode == "1" {
		// log to json format for friendly with logstash/graylog/etc...
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	log.SetOutput(os.Stdout)
}
