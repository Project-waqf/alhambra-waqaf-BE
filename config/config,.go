package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	SERVER_PORT int
	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     int
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	SALT1       string
	SALT2       string
	SECRET_JWT  string
	REDIS_HOST  string
	REDIS_PORT  string
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

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("ERROR: load .env", err)
		}
	}

	serverPort, errPortServer := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if errPortServer != nil {
		log.Fatal("Error parse SERVER_PORT", errPortServer)
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
	defaultConfig.SALT1 = os.Getenv("SALT1")
	defaultConfig.SALT2 = os.Getenv("SALT2")
	defaultConfig.REDIS_HOST = os.Getenv("REDIS_HOST")
	defaultConfig.REDIS_PORT = os.Getenv("REDIS_PORT")

	return &defaultConfig
}
