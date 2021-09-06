package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
	// "os/user"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}
