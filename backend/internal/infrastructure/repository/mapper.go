package repository

import (
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/infraModels"
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
	//	TODO: Implement this function
	panic("function not implemented")
}

func toInfraPlans(plans []models.Plan) []infraModels.Plan {
	//	TODO: Implement this function
	panic("function not implemented")
}

func toDomainUserWeightHistory(histories []infraModels.UserWeightHistory) []models.UserWeightHistory {
	//	TODO: Implement this function
	panic("function not implemented")
}

func toInfraUserWeightHistory(histories []models.UserWeightHistory) []infraModels.UserWeightHistory {
	//	TODO: Implement this function
	panic("function not implemented")
}
