package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	// "github.com/joho/godotenv"
)

type AppConfig struct {
	SERVER_PORT int
	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	SECRET_JWT  string
}

var lock = &sync.Mutex{}
var config *AppConfig

func Getconfig() *AppConfig {
	defer lock.Unlock()
	lock.Lock()

	if config == nil {
		config = initConfig()
	}

	return config
}

func initConfig() *AppConfig {

	var defaultConfig AppConfig

	// if err := godotenv.Load("config.env"); err != nil {
	// 	log.Fatal(err)
	// }

	serverPort, errPortServer := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if errPortServer != nil {
		log.Fatal("Error parse SERVER_PORT")
	}
	dbPort, errPortDB := strconv.Atoi(os.Getenv("DB_PORT"))
	if errPortDB != nil {
		log.Fatal("Error parse DB_PORT")
	}
	defaultConfig.DB_PORT = dbPort
	defaultConfig.SERVER_PORT = serverPort
	defaultConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.SECRET_JWT = os.Getenv("SECRET_JWT")

	return &defaultConfig
}
