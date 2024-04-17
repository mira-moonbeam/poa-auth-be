package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mira-moonbeam/go-auth-be/utils/config"
	"log"
)

var DB *gorm.DB

var configMap config.Config

func init() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	configMap = loadConfig
}

func ConnectDatabase() {
	driver := configMap.DBDriver
	host := configMap.DBHost
	user := configMap.DBUser
	password := configMap.DBPassword
	name := configMap.DBName
	port := configMap.DBPort

	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, user, password, host, port, name)
	var err error
	DB, err = gorm.Open(driver, dbUrl)

	if err != nil {
		fmt.Println("Cannot connect to database ", driver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database", driver)
	}

	DB.AutoMigrate(&User{})
}
