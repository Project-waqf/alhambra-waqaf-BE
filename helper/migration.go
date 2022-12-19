package helper

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&Admin{})
}
