package models

type Gender string
type Goal string
type PhysicalActivity string

const (
	Male           Gender           = "Male"
	Female         Gender           = "Female"
	LoseWeight     Goal             = "LoseWeight"
	GainMuscleMass Goal             = "GainMuscleMass"
	Low            PhysicalActivity = "Low"
	Medium         PhysicalActivity = "Medium"
	High           PhysicalActivity = "High"
)

type User struct {
	Utgid             int64               `gorm:"column:utgid;primaryKey"`
	Name              string              `gorm:"column:name"`
	Gender            Gender              `gorm:"column:gender"`
	Weight            int                 `gorm:"column:weight"`
	Height            int                 `gorm:"column:height"`
	Goal              Goal                `gorm:"column:goal"`
	PhysicalActivity  PhysicalActivity    `gorm:"column:physical_activity"`
	Actions           []Action            `gorm:"foreignKey:Utgid;references:Utgid"`
	Plans             []Plan              `gorm:"foreignKey:Utgid;references:Utgid"`
	UserWeightHistory []UserWeightHistory `gorm:"foreignKey:Utgid;references:Utgid"`
}
