package repository

import (
	"fmt"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Action{},
		&models.Plan{},
		&models.UserWeightHistory{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
