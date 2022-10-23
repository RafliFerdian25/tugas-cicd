package mock

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) GetAllUsers() ([]dto.UserDTO, error) {
	args := u.Called()

	return args.Get(0).([]dto.UserDTO), args.Error(1)
}

func (u *UserMock) CreateUser(data model.User) error {
	args := u.Called(data)

	return args.Error(0)
}