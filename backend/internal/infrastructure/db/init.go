package db

import (
	"fmt"
	"github.com/isiyar/daily-energy/backend/config"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/infraModels"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&infraModels.User{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&infraModels.Action{},
		&infraModels.Plan{},
		&infraModels.UserWeightHistory{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
