package repository

import (
	"context"
	"fmt"

	"github.com/isiyar/daily-energy/backend/internal/adapters/adapterModels"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"gorm.io/gorm"
)

type userWeightHistoryRepository struct {
	db *gorm.DB
}

func NewUserWeightHistoryRepository(db *gorm.DB) ports.UserWeightHistoryRepository {
	return &userWeightHistoryRepository{db: db}
}

func (r *userWeightHistoryRepository) Save(ctx context.Context, userWeightHistory models.UserWeightHistory) error {
	userWeightHistoryArray := []models.UserWeightHistory{userWeightHistory}
	userWeightHistoryAdapter := toAdapterUserWeightHistory(userWeightHistoryArray)[0]
	if err := r.db.WithContext(ctx).Save(&userWeightHistoryAdapter).Error; err != nil {
		return err
	}
	return nil
}

func (r *userWeightHistoryRepository) GetUserWeightHistory(ctx context.Context, utgid int64) ([]models.UserWeightHistory, error) {
	var adapterUserWeightHistory []adapterModels.UserWeightHistory
  
  //   err := r.db.WithContext(ctx).
  //     Where("utgid", utgid,).
//     Find(&adapterUserWeightHistory).Error

//   if err != nil {
//     return nil, err
//   }
	if err := r.db.WithContext(ctx).Where("utgid = ?", utgid).Find(&adapterUserWeightHistory).Error; err != nil {
		return []models.UserWeightHistory{}, err
	}
	fmt.Println()

  	return toDomainUserWeightHistory(adapterUserWeightHistory), nil
}
