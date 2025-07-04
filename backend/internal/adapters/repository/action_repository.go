package repository

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/adapters/adapterModels"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"gorm.io/gorm"
)

type actionRepository struct {
	db *gorm.DB
}

func NewActionRepository(db *gorm.DB) ports.ActionRepository {
	return &actionRepository{db: db}
}

func (r *actionRepository) GetById(ctx context.Context, id string) (models.Action, error) {
	var a adapterModels.Action
	if err := r.db.WithContext(ctx).First(&a, "id = ?", id).Error; err != nil {
		return models.Action{}, err
	}
	return toDomainAction(a), nil
}

func (r *actionRepository) GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishtAt int) (models.Action, error) {
	//TODO implement me
	panic("implement me")
}

func (r *actionRepository) Save(ctx context.Context, user models.Action) error {
	//TODO implement me
	panic("implement me")
}

func (r *actionRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
