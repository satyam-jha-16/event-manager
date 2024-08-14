package repositories

import (
	"context"

	"github.com/satyam-jha-16/event-manager/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) RegisterUser(context context.Context, registerData *models.AuthCredentials) (*models.User, error) {
	user := &models.User{
		Email:    registerData.Email,
		Password: registerData.Password,
	}
	res := r.db.Model(&models.User{}).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {
	user := &models.User{}

	if res := r.db.Model(user).Where(query, args...).First(user); res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
