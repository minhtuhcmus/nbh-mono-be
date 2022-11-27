package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type UserRepository struct{}

var userRepository *UserRepository

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (u *UserRepository) GetUserByUsername(
	ctx context.Context,
	username string,
	userData *models.User,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Where("username = ?", username).
		Where("active").
		First(userData).Error
	if err != nil {
		return fmt.Errorf("error UserRepository.GetUserByUsername: %v", err)
	}
	return nil
}

func (u *UserRepository) GetRoleByUserID(
	ctx context.Context,
	userID int,
	roles *[]string,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT r.label FROM user_roles ur INNER JOIN roles r ON ur.fk_role = r.id WHERE ur.fk_user = ?", userID).
		Scan(&roles).Error
	if err != nil {
		return fmt.Errorf("error UserRepository.GetRoleByUserID %v", err)
	}
	return nil
}
