package services

import (
	"resttest/domain/users"
	"resttest/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
