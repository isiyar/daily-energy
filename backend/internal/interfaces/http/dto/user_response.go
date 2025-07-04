package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type UserResponse struct {
	Utgid            int64                   `json:"utgid"`
	Name             string                  `json:"name" validate:"required,min=1,max=50"`
	Gender           models.Gender           `json:"gender"`
	Weight           int                     `json:"weight" validate:"required,min=10,max=300"`
	Height           int                     `json:"height" validate:"required,min=10,max=300"`
	Goal             models.Goal             `json:"goal"`
	PhysicalActivity models.PhysicalActivity `json:"physical_activity"`
}

func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		Utgid:            u.Utgid,
		Name:             u.Name,
		Gender:           u.Gender,
		Weight:           u.Weight,
		Height:           u.Height,
		Goal:             u.Goal,
		PhysicalActivity: u.PhysicalActivity,
	}
}
