package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/afkpanda2004/golang-todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,

) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}

}

func NewUserUnitialized(
	fullName string,
	phoneNumber *string,

) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumber := len([]rune(*u.PhoneNumber))
		if phoneNumber < 10 || phoneNumber > 15 {
			return fmt.Errorf(
				"invalid `FullName` len: %d: %w",
				phoneNumber,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `FullName` len: %w",
				core_errors.ErrInvalidArgument,
			)

		}

	}

	return nil
}
