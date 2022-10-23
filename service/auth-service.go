package service

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type AuthService interface {
	LoginUser(user model.User) (model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo,
	}
}

func (a *authService) LoginUser(user model.User) (model.User, error) {
	user, err := a.userRepo.LoginUser(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}