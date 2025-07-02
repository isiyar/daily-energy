package models

type ActionType string

const (
	Food     ActionType = "Food"
	Activity ActionType = "Activity"
)

type Action struct {
	Utgid        int64      `gorm:"column:utgid"`
	Date         int64      `gorm:"column:date"`
	ActivityName string     `gorm:"column:activity_name"`
	Calories     int        `gorm:"column:calories"`
	Type         ActionType `gorm:"column:type"`
}
