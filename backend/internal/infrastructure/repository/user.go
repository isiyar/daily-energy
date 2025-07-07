package repository

import (
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type UserRepository interface {
	Create() models.User
	Read() models.User
}
