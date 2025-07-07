package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type ActionRequest struct {
	Date         int64             `json:"date" validate:"required"`
	ActivityName string            `json:"activity_name" validate:"required"`
	Calories     int               `json:"calories" validate:"required,min=1"`
	Type         models.ActionType `json:"type" validate:"required"`
}

func (a *ActionRequest) ToAction(utgid int64) models.Action {
	return models.Action{
		Utgid:        utgid,
		Date:         a.Date,
		ActivityName: a.ActivityName,
		Calories:     a.Calories,
		Type:         a.Type,
	}
}
