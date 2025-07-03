package repository

import (
	"context"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
	"github.com/isiyar/daily-energy/backend/internal/domain/ports"
	"github.com/isiyar/daily-energy/backend/internal/infrastructure/infraModels"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUtgid(ctx context.Context, utgid int64) (models.User, error) {
	var u infraModels.User
	if err := r.db.WithContext(ctx).First(&u, "utgid = ?", utgid).Error; err != nil {
		return models.User{}, err
	}
	return toDomainUser(u), nil
}

func (r *userRepository) Save(ctx context.Context, user models.User) error {
	userInfra := toInfraUser(user)
	if err := r.db.WithContext(ctx).Save(&userInfra).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, utgid int64) error {
	if err := r.db.WithContext(ctx).Delete(&infraModels.User{}, "utgid = ?", utgid).Error; err != nil {
		return err
	}
	return nil
}
