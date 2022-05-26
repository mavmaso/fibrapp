package db

import (
	"fibrapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Article{})
}
