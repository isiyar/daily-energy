package repository

import (
	"context"
	"errors"
	"github.com/isiyar/daily-energy/backend/internal/adapters/adapterModels"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"github.com/isiyar/daily-energy/backend/internal/interfaces/http/dto"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUtgid(ctx context.Context, utgid int64) (models.User, error) {
	var u adapterModels.User
	if err := r.db.WithContext(ctx).First(&u, "utgid = ?", utgid).Error; err != nil {
		return models.User{}, err
	}
	return toDomainUser(u), nil
}

func (r *userRepository) Save(ctx context.Context, user models.User) error {
	userAdapter := toAdapterUser(user)
	if err := r.db.WithContext(ctx).Save(&userAdapter).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, utgid int64) error {
	if err := r.db.WithContext(ctx).Delete(&adapterModels.User{}, "utgid = ?", utgid).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(
	ctx context.Context, utgid int64, req dto.UserRequest) (models.User, error) {
	params := make(map[string]interface{})
	if req.Gender != "" {
		params["gender"] = req.Gender
	}
	if req.Name != "" {
		params["name"] = req.Name
	}
	if req.Goal != "" {
		params["goal"] = req.Goal
	}
	if req.PhysicalActivity != "" {
		params["physical_activity"] = req.PhysicalActivity
	}
	if req.Weight > 0 {
		params["weight"] = req.Weight
	}
	if req.Height > 0 {
		params["height"] = req.Height
	}
	if len(params) == 0 {
		return models.User{}, errors.New("no fields to update")
	}

	var new_user adapterModels.User
	if err := r.db.WithContext(ctx).Model(&adapterModels.User{}).Where("utgid = ?", utgid).Updates(params).First(&new_user).Error; err != nil {
		return models.User{}, err
	}
	return toDomainUser(new_user), nil
}
