package services

import (
	"github.com/alvaro259818/bookstore_users-api/domain/users"
	"github.com/alvaro259818/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil

	return nil, &errors.RestError{
		Status: http.StatusInternalServerError,
	}
}