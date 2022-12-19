package helper

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}