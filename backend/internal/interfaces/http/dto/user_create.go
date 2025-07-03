package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type UserCreate struct {
	Utgid            int64                   `json:"utgid" validate:"required,min=1"`
	Name             string                  `json:"name" validate:"required,min=1,max=50"`
	Gender           models.Gender           `json:"gender"`
	Weight           int                     `json:"weight" validate:"required,min=10,max=300"`
	Height           int                     `json:"height" validate:"required,min=10,max=300"`
	Goal             models.Goal             `json:"goal"`
	PhysicalActivity models.PhysicalActivity `json:"physical_activity"`
}

func (u *UserCreate) ToUser() models.User {
	return models.User{
		Utgid:            u.Utgid,
		Name:             u.Name,
		Gender:           u.Gender,
		Weight:           u.Weight,
		Height:           u.Height,
		Goal:             u.Goal,
		PhysicalActivity: u.PhysicalActivity,
	}
}
