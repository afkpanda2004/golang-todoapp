package users_service

import (
	"context"

	"github.com/afkpanda2004/golang-todoapp/internal/core/domain"
)

type UsersService struct {
	userRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,

	) (domain.User, error)
}

func NewUsersService(
	userRepository UsersRepository,
) *UsersService {
	return &UsersService{
		userRepository: userRepository,
	}
}
