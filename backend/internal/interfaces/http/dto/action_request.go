package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type ActionRequest struct {
	Date         int64             `json:"date" validate:"required"`
	ActivityName string            `json:"activity_name" validate:"required"`
	Calories     int               `json:"calories" validate:"required,min=1"`
	Type         models.ActionType `json:"type" validate:"required"`
}
