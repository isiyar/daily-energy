package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/isiyar/daily-energy/backend/internal/adapters/adapterModels"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"github.com/isiyar/daily-energy/backend/pkg/utils"
	"gorm.io/gorm"
)

type actionRepository struct {
	db *gorm.DB
}

func NewActionRepository(db *gorm.DB) ports.ActionRepository {
	return &actionRepository{db: db}
}

func (r *actionRepository) GetById(ctx context.Context, id string) (models.Action, error) {
	aid, err := uuid.Parse(id)
	if err != nil {
		return models.Action{}, err
	}

	var a adapterModels.Action
	if err := r.db.WithContext(ctx).First(&a, "id = ?", aid).Error; err != nil {
		return models.Action{}, err
	}
	return toDomainAction(a), nil
}

func (r *actionRepository) GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishtAt int) (models.Action, error) {
	//TODO implement me
	panic("implement me")
}

func (r *actionRepository) Save(ctx context.Context, action *models.Action) error {
	action.Id = utils.GenerateUUID().String()

	actionAdapter := toAdapterAction(*action)
	if err := r.db.WithContext(ctx).Save(&actionAdapter).Error; err != nil {
		return err
	}
	return nil
}
