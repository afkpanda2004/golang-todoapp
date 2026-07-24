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
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,

	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
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
