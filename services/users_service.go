package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/crypto_utils"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"log"
	// "os/user"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	log.Println(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	currentUser := &users.User{Id: user.Id}

	if err := currentUser.Get(); err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}
	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{
		Id: userId,
	}

	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)

}
