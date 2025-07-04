package repository

import (
	"github.com/isiyar/daily-energy/backend/internal/adapters/infraModels"
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

func toDomainUser(u infraModels.User) models.User {
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

func toInfraUser(u models.User) infraModels.User {
	return infraModels.User{
		Utgid:             u.Utgid,
		Name:              u.Name,
		Gender:            u.Gender,
		Weight:            u.Weight,
		Height:            u.Height,
		Goal:              u.Goal,
		PhysicalActivity:  u.PhysicalActivity,
		Actions:           toInfraActions(u.Actions),
		Plans:             toInfraPlans(u.Plans),
		UserWeightHistory: toInfraUserWeightHistory(u.UserWeightHistory),
	}
}

func toDomainActions(actions []infraModels.Action) []models.Action {
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

func toInfraActions(actions []models.Action) []infraModels.Action {
	if actions == nil {
		return nil
	}
	res := make([]infraModels.Action, len(actions))
	for i, a := range actions {
		id, err := ParseUUID(a.Id)
		if err != nil {
			continue
		}
		res[i] = infraModels.Action{
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

func toDomainPlans(plans []infraModels.Plan) []models.Plan {
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

func toInfraPlans(plans []models.Plan) []infraModels.Plan {
	if plans == nil {
		return nil
	}
	res := make([]infraModels.Plan, len(plans))
	for i, p := range plans {
		id, err := ParseUUID(p.Id)
		if err != nil {
			continue
		}
		res[i] = infraModels.Plan{
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

func toDomainUserWeightHistory(histories []infraModels.UserWeightHistory) []models.UserWeightHistory {
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

func toInfraUserWeightHistory(histories []models.UserWeightHistory) []infraModels.UserWeightHistory {
	if histories == nil {
		return nil
	}
	res := make([]infraModels.UserWeightHistory, len(histories))
	for i, h := range histories {
		id, err := ParseUUID(h.Id)
		if err != nil {
			continue
		}
		res[i] = infraModels.UserWeightHistory{
			Id:         id,
			Utgid:      h.Utgid,
			Date:       h.Date,
			UserWeight: h.UserWeight,
			UserHeight: h.UserHeight,
		}
	}
	return res
}
