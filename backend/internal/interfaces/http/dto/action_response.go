package dto

import "github.com/isiyar/daily-energy/backend/internal/domain/models"

type ActionResponse struct {
	Id           string            `json:"id"`
	Utgid        int64             `json:"utgid"`
	Date         int64             `json:"date"`
	ActivityName string            `json:"activity_name"`
	Calories     int               `json:"calories"`
	Type         models.ActionType `json:"type"`
}

func ToActionResponse(a models.Action) ActionResponse {
	return ActionResponse{
		Id:           a.Id,
		Utgid:        a.Utgid,
		Date:         a.Date,
		ActivityName: a.ActivityName,
		Calories:     a.Calories,
		Type:         a.Type,
	}
}
