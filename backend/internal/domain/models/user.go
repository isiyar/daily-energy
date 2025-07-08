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
	Utgid             int64
	Name              string
	Gender            Gender
	DateofBirth       int64
	Weight            int
	Height            int
	Goal              Goal
	PhysicalActivity  PhysicalActivity
	Actions           []Action
	Plans             []Plan
	UserWeightHistory []UserWeightHistory
}
