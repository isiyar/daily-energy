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
