package mock

import (
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (a *AuthMock) LoginUser(user model.User) (model.User, error) {
	args := a.Called(user)

	return args.Get(0).(model.User), args.Error(1)
}
