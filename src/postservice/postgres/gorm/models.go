package gorm

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Post struct {
	ID          string `gorm:"primarykey"`
	title       string
	description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func InitialMigrationUserProfile(db *gorm.DB) {
	if err := db.AutoMigrate(&Post{}); err != nil {
		log.Fatalln(err)
	}
	//Check if table exist or not
	if !db.Migrator().HasTable(&Post{}) {
		// if table not exist then create table
		if err := db.Migrator().CreateTable(&Post{}); err != nil {
			log.Fatalln(err)
		}
	}
}
