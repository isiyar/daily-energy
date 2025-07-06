package adapterModels

import (
	// "database/sql/driver"
	// "fmt"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type User struct {
	Utgid             int64                   `gorm:"column:utgid;primaryKey"`
	Name              string                  `gorm:"column:name"`
	Gender            models.Gender           `gorm:"column:gender"`
	Weight            int                     `gorm:"column:weight"`
	Height            int                     `gorm:"column:height"`
	Goal              models.Goal             `gorm:"column:goal"`
	PhysicalActivity  models.PhysicalActivity `gorm:"column:physical_activity"`
	Actions           []Action                `gorm:"foreignKey:Utgid"`
	Plans             []Plan                  `gorm:"foreignKey:Utgid"`
	UserWeightHistory []UserWeightHistory     `gorm:"foreignKey:Utgid"`
}

// // Gender implementation
// type Gender string

// const (
// 	Male   Gender = "Male"
// 	Female Gender = "Female"
// )

// func (g *Gender) Scan(value interface{}) error {
// 	if value == nil {
// 		return nil
// 	}
// 	switch v := value.(type) {
// 	case []byte:
// 		*g = Gender(v)
// 	case string:
// 		*g = Gender(v)
// 	default:
// 		return fmt.Errorf("unsupported type for Gender: %T", value)
// 	}
// 	return nil
// }

// func (g Gender) Value() (driver.Value, error) {
// 	return string(g), nil
// }

// // Goal implementation
// type Goal string

// const (
// 	WeightLoss   Goal = "WeightLoss"
// 	WeightGain   Goal = "WeightGain"
// 	Maintenance  Goal = "Maintenance"
// 	MuscleGain   Goal = "MuscleGain"
// 	Endurance    Goal = "Endurance"
// )

// func (gl *Goal) Scan(value interface{}) error {
// 	if value == nil {
// 		return nil
// 	}
// 	switch v := value.(type) {
// 	case []byte:
// 		*gl = Goal(v)
// 	case string:
// 		*gl = Goal(v)
// 	default:
// 		return fmt.Errorf("unsupported type for Goal: %T", value)
// 	}
// 	return nil
// }

// func (gl Goal) Value() (driver.Value, error) {
// 	return string(gl), nil
// }

// // PhysicalActivity implementation
// type PhysicalActivity string

// const (
// 	Sedentary    PhysicalActivity = "Sedentary"
// 	LightActive  PhysicalActivity = "LightActive"
// 	Moderate     PhysicalActivity = "Moderate"
// 	VeryActive   PhysicalActivity = "VeryActive"
// 	Extreme      PhysicalActivity = "Extreme"
// )

// func (pa *PhysicalActivity) Scan(value interface{}) error {
// 	if value == nil {
// 		return nil
// 	}
// 	switch v := value.(type) {
// 	case []byte:
// 		*pa = PhysicalActivity(v)
// 	case string:
// 		*pa = PhysicalActivity(v)
// 	default:
// 		return fmt.Errorf("unsupported type for PhysicalActivity: %T", value)
// 	}
// 	return nil
// }

// func (pa PhysicalActivity) Value() (driver.Value, error) {
// 	return string(pa), nil
// }
