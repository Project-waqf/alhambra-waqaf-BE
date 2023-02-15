package helper

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255) type:not null"`
	Email    string `gorm:"type:varchar(255) type:not null"`
	Password string `gorm:"type:varchar(255) type:not null"`
}

type News struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255) type:not null"`
	Body    string `gorm:"type:text type:not null"`
	Picture string `gorm:"type:varchar(255) type:not null"`
	Status  string `gorm:"type:enum('draft', 'online', 'archive')"`
}

type Wakaf struct {
	gorm.Model
	Title      string `gorm:"type:varchar(255) type:not null"`
	Category   string `gorm:"type:varchar(255) type:not null"`
	Picture    string `gorm:"type:varchar(255) type:not null"`
	Collected  int    `gorm:"type:int(11) type:not null"`
	FundTarget int    `gorm:"type:int(11) type:not null"`
	DateTarget *time.Time 
}

type Asset struct {
	gorm.Model
	Name    string `gorm:"varchar(255) type:not null"`
	Picture string `gorm:"varchar(255) type:not null"`
	Detail  string `gorm:"varchar(255) type:not null"`
	Status  string `gorm:"type:enum('draft', 'online', 'archive')"`
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&News{})
	db.AutoMigrate(&Wakaf{})
	db.AutoMigrate(&Asset{})
}
