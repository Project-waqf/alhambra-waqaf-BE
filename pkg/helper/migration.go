package helper

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type News struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Body    string `gorm:"type:longtext;not null"`
	Picture string `gorm:"type:varchar(255);not null"`
	Status  string `gorm:"type:enum('draft', 'online', 'archive')"`
	FileId  string `gorm:"type:varchar(255)"`
}

type Wakaf struct {
	gorm.Model
	Title      string     `gorm:"type:varchar(255);not null"`
	Detail     string     `gorm:"type:longtext;not null"`
	Category   string     `gorm:"type:enum('pendidikan', 'bangunan', 'umum', 'kesehatan');not null"`
	Picture    string     `gorm:"type:varchar(255);not null"`
	Collected  int        `gorm:"not null"`
	FundTarget int        `gorm:"not null"`
	DueDate    *time.Time `gorm:"type:datetime;not null"`
	FileId     string     `gorm:"type:varchar(255)"`
	Donor      []Donor    `gorm:"foreignKey:IdWakaf"`
}

type Asset struct {
	gorm.Model
	Name    string `gorm:"varchar(255);not null"`
	Picture string `gorm:"varchar(255);not null"`
	Detail  string `gorm:"longtext;not null"`
	Status  string `gorm:"type:enum('draft', 'online', 'archive')"`
	FileId  string `gorm:"type:varchar(255)"`
}

type Donor struct {
	gorm.Model
	IdWakaf     uint
	Name        string `gorm:"type:varchar(255);not null"`
	Doa         string `gorm:"type:varchar(255)"`
	GrossAmount string `gorm:"type:int(13);not null"`
	Status      string `gorm:"type:varchar(20)"`
	OrderId     string `gorm:"type:varchar(255)"`
	PaymentType string `gorm:"type:varchar(255)"`
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&News{})
	db.AutoMigrate(&Wakaf{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(Donor{})
}
