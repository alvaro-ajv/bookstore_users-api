package services

import (
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"github.com/alvaro259818/bookstore_users-api/domain/users"
	"github.com/alvaro259818/bookstore_users-api/utils/crypto_utils"
	"github.com/alvaro259818/bookstore_users-api/utils/date_utils"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *rest_errors.RestError)
	CreateUser(users.User) (*users.User, *rest_errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestError)
	DeleteUser(int64) *rest_errors.RestError
	SearchUser(string) (users.Users, *rest_errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestError)
}

func (s *usersService) GetUser(userId int64) (*users.User, *rest_errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestError) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *rest_errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *rest_errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *rest_errors.RestError)  {
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}