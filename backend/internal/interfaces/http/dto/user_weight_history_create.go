package dto

import (
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/pkg"
)

type UserWeightHistoryCreate struct {
	Date   int64 `json:"date" validate:"required"`
	Weight int   `json:"weight" validate:"required,min=10,max=300"`
	Height int   `json:"height" validate:"required,min=10,max=300"`
}

func (uw *UserWeightHistoryCreate) ToUserWeightHistory(utgid int64) models.UserWeightHistory {
	return models.UserWeightHistory{
		Id:         utils.GenerateUUID().String(),
		Utgid:      utgid,
		Date:       uw.Date,
		UserWeight: uw.Weight,
		UserHeight: uw.Height,
	}
}
