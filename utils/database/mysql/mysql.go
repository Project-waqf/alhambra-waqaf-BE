package mysql

import (
	"fmt"
	"log"
	"wakaf/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBmysql(config *config.AppConfig) *gorm.DB {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error initialize connection MySQL")
	}
	return db
}