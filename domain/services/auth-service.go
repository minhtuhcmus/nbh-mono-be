package services

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/utils"
)

type AuthService interface {
	SignIn(ctx context.Context, username, password string) (string, error)
}

type authService struct {
	userRepository *repositories.UserRepository
}

func (a authService) SignIn(ctx context.Context, username, password string) (string, error) {
	var userDetail models.User
	err := a.userRepository.GetUserByUsername(ctx, username, &userDetail)
	if err != nil {
		return "", err
	}
	isMatch := utils.CheckPasswordHash(password, userDetail.HashedPassword)
	if isMatch {
		var roles []string
		err = a.userRepository.GetRoleByUserID(ctx, userDetail.ID, &roles)
		if err != nil {
			return "", err
		}
		return utils.GenTokenPair(userDetail.ID, roles)
	} else {
		return "", fmt.Errorf("error authService.SignIn username and password not match")
	}
}

func NewAuthService(
	userRepository *repositories.UserRepository,
) AuthService {
	return &authService{userRepository: userRepository}
}
