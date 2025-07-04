package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type UserRequest struct {
	Name             string                  `json:"name" validate:"required,min=1,max=50"`
	Gender           models.Gender           `json:"gender" validate:"required"`
	Goal             models.Goal             `json:"goal" validate:"required"`
	PhysicalActivity models.PhysicalActivity `json:"physical_activity" validate:"required"`
}
