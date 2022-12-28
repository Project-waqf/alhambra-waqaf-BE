package helper

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

type News struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255)"`
	Body    string `gorm:"type:text"`
	Picture string `gorm:"type:varchar(255)"`
}

type Wakaf struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255)"`
	Category string `gorm:"type:varchar(255)"`
	Picture  string `gorm:"type:varchar(255)"`
}

type Asset struct {
	gorm.Model
	Name    string `gorm:"varchar(255)"`
	Picture string `gorm:"varchar(255)"`
	Detail  string `gorm:"varchar(255)"`
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&News{})
	db.AutoMigrate(&Wakaf{})
	db.AutoMigrate(&Asset{})
}
