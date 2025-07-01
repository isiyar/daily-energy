package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.DB_HOST, c.DB_PORT, c.DB_USERNAME, c.DB_PASSWORD, c.DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// После написания моделек закинуть сюда
	// db.AutoMigrate()

	return db, nil
}
