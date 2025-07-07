package adapterModels

import "github.com/google/uuid"

type Plan struct {
	Id                uuid.UUID `gorm:"column:id;primaryKey"`
	Utgid             int64       `gorm:"column:utgid"`
	Date              int64       `gorm:"column:date"`
	CaloriesToConsume int         `gorm:"column:calories_to_consume"`
	CaloriesToBurn    int         `gorm:"column:calories_to_burn"`
	Recommendation    string      `gorm:"column:recommendation"`
}
