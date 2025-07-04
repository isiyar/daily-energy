package repository

import (
	"github.com/isiyar/daily-energy/backend/internal/adapters/adapterModels"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func ParseUUID(s string) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(s)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func toDomainUser(u adapterModels.User) models.User {
	return models.User{
		Utgid:             u.Utgid,
		Name:              u.Name,
		Gender:            u.Gender,
		Weight:            u.Weight,
		Height:            u.Height,
		Goal:              u.Goal,
		PhysicalActivity:  u.PhysicalActivity,
		Actions:           toDomainActions(u.Actions),
		Plans:             toDomainPlans(u.Plans),
		UserWeightHistory: toDomainUserWeightHistory(u.UserWeightHistory),
	}
}

func toAdapterUser(u models.User) adapterModels.User {
	return adapterModels.User{
		Utgid:             u.Utgid,
		Name:              u.Name,
		Gender:            u.Gender,
		Weight:            u.Weight,
		Height:            u.Height,
		Goal:              u.Goal,
		PhysicalActivity:  u.PhysicalActivity,
		Actions:           toAdapterActions(u.Actions),
		Plans:             toAdapterPlans(u.Plans),
		UserWeightHistory: toAdapterUserWeightHistory(u.UserWeightHistory),
	}
}

func toDomainActions(actions []adapterModels.Action) []models.Action {
	if actions == nil {
		return nil
	}
	res := make([]models.Action, len(actions))
	for i, a := range actions {
		res[i] = models.Action{
			Id:           a.Id.String(),
			Utgid:        a.Utgid,
			Date:         a.Date,
			ActivityName: a.ActivityName,
			Calories:     a.Calories,
			Type:         a.Type,
		}
	}
	return res
}

func toAdapterActions(actions []models.Action) []adapterModels.Action {
	if actions == nil {
		return nil
	}
	res := make([]adapterModels.Action, len(actions))
	for i, a := range actions {
		id, err := ParseUUID(a.Id)
		if err != nil {
			continue
		}
		res[i] = adapterModels.Action{
			Id:           id,
			Utgid:        a.Utgid,
			Date:         a.Date,
			ActivityName: a.ActivityName,
			Calories:     a.Calories,
			Type:         a.Type,
		}
	}
	return res
}

func toDomainPlans(plans []adapterModels.Plan) []models.Plan {
	if plans == nil {
		return nil
	}
	res := make([]models.Plan, len(plans))
	for i, p := range plans {
		res[i] = models.Plan{
			Id:                p.Id.String(),
			Utgid:             p.Utgid,
			Date:              p.Date,
			CaloriesToConsume: p.CaloriesToConsume,
			CaloriesToBurn:    p.CaloriesToBurn,
			Recommendation:    p.Recommendation,
		}
	}
	return res
}

func toAdapterPlans(plans []models.Plan) []adapterModels.Plan {
	if plans == nil {
		return nil
	}
	res := make([]adapterModels.Plan, len(plans))
	for i, p := range plans {
		id, err := ParseUUID(p.Id)
		if err != nil {
			continue
		}
		res[i] = adapterModels.Plan{
			Id:                id,
			Utgid:             p.Utgid,
			Date:              p.Date,
			CaloriesToConsume: p.CaloriesToConsume,
			CaloriesToBurn:    p.CaloriesToBurn,
			Recommendation:    p.Recommendation,
		}
	}
	return res
}

func toDomainUserWeightHistory(histories []adapterModels.UserWeightHistory) []models.UserWeightHistory {
	if histories == nil {
		return nil
	}
	res := make([]models.UserWeightHistory, len(histories))
	for i, h := range histories {
		res[i] = models.UserWeightHistory{
			Id:         h.Id.String(),
			Utgid:      h.Utgid,
			Date:       h.Date,
			UserWeight: h.UserWeight,
			UserHeight: h.UserHeight,
		}
	}
	return res
}

func toAdapterUserWeightHistory(histories []models.UserWeightHistory) []adapterModels.UserWeightHistory {
	if histories == nil {
		return nil
	}
	res := make([]adapterModels.UserWeightHistory, len(histories))
	for i, h := range histories {
		id, err := ParseUUID(h.Id)
		if err != nil {
			continue
		}
		res[i] = adapterModels.UserWeightHistory{
			Id:         id,
			Utgid:      h.Utgid,
			Date:       h.Date,
			UserWeight: h.UserWeight,
			UserHeight: h.UserHeight,
		}
	}
	return res
}
