package gorm

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	UserName    string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func InitialMigrationUserProfile(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalln(err)
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&User{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&User{}); err != nil {
			log.Fatalln(err)
		}
	}
}
