package models

type UserWeightHistory struct {
	Utgid      int64 `gorm:"column:utgid"`
	Date       int64 `gorm:"column:date"`
	UserWeight int   `gorm:"column:user_weight"`
	UserHeight int   `gorm:"column:user_height"`
}
