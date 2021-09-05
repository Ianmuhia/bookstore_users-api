package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
	// "os/user"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}
