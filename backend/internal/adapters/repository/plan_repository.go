package repository

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"gorm.io/gorm"
)

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) ports.PlanRepository {
	return &planRepository{db: db}
}

func (r *planRepository) GetByStartTimeAndFinishTime(ctx context.Context, StartAt, FinishAt int64) (models.Plan, error) {
	//TODO implement me
	panic("implement me")
}

func (r *planRepository) Save(ctx context.Context, user models.Plan) error {
	//TODO implement me
	panic("implement me")
}

func (r *planRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
