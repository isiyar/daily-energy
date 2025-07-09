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

func (r *planRepository) GetByStartTimeAndFinishTime(ctx context.Context, startAt, finishAt, utgid int64) ([]models.Plan, error) {
	//TODO implement me
	panic("implement me")
}

func (r *planRepository) Save(ctx context.Context, plans []models.Plan) error {
	plansAdapter := toAdapterPlans(plans)
	if err := r.db.WithContext(ctx).Save(&plansAdapter).Error; err != nil {
		return err
	}
	return nil
}
