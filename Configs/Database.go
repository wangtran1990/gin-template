package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

// DB ... DBConfig represents db configuration
var DB *gorm.DB
var err error

// Config variables from .env
var dbstatus, dbhost, dbuser, dbpassword, dbname string
var dbport int

// InitDB ...
func InitDB() error {
	dbstatus = os.Getenv("RDB_ON")
	dbhost = os.Getenv("RDB_HOST")
	dbport, _ = strconv.Atoi(os.Getenv("RDB_PORT"))
	dbuser = os.Getenv("RDB_USER")
	dbpassword = os.Getenv("RDB_PASSWORD")
	dbname = os.Getenv("RDB_DBNAME")

	if dbstatus != "1" { // ON
		return errors.New("RDB off by setting")
	}

	DB, err = gorm.Open("mysql", _dbURL(_buildDBConfig()))
	if err != nil {
		return err
	}

	return nil
}

/************************************************************/
// PRIVATE Functions                                        */
/************************************************************/

type _dbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func _buildDBConfig() *_dbConfig {
	dbConfig := _dbConfig{
		Host:     dbhost,
		Port:     dbport,
		User:     dbuser,
		Password: dbpassword,
		DBName:   dbname,
	}
	return &dbConfig
}

func _dbURL(dbConfig *_dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
