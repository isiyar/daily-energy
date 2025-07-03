package models

import "github.com/jackc/pgx/v5/pgtype"

type UserWeightHistory struct {
	Id         pgtype.UUID `gorm:"column:id;primaryKey"`
	Utgid      int64       `gorm:"column:utgid"`
	Date       int64       `gorm:"column:date"`
	UserWeight int         `gorm:"column:user_weight"`
	UserHeight int         `gorm:"column:user_height"`
}
