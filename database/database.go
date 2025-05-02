package database

import (
	"reuse-api/config"
	models "reuse-api/models/user"

	items "reuse-api/models/items"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.GetEnv("DB_DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar com o banco")
	}
	DB = database
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&items.Item{})
}
