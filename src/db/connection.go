package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type envVars struct {
	dbType string
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

var config envVars
var db *gorm.DB

func init() {
	loadVars()
	connectionString := config.dbUser + ":" + config.dbPass + "@(" + config.dbHost + ":" + config.dbPort + ")/" + config.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(config.dbType, connectionString)
	if err != nil {
		log.Fatal("Failed to connect database, ", err)
	}

	db = conn
	log.Println("DB Connected")

	err = db.Debug().DropTableIfExists(&Rates{}).Error // This is for no repeat information
	if err != nil {
		log.Fatal("Failed to drop table, ", err)
	}
	err = db.Debug().AutoMigrate(&Rates{}).Error
	if err != nil {
		log.Fatal("Failed to migrate table, ", err)
	}
}

func loadVars() {
	valuesEnv := godotenv.Load()
	if valuesEnv != nil {
		log.Fatal("Error loading environment variables")
	}
	config = envVars{
		dbType: os.Getenv("DB_TYPE"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASS"),
		dbHost: os.Getenv("DB_HOST"),
		dbPort: os.Getenv("DB_PORT"),
		dbName: os.Getenv("DB_NAME"),
	}

}
