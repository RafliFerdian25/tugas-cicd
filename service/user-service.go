package service

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserService interface {
	GetAllUsers() ([]dto.UserDTO, error)
	CreateUser(user model.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return []dto.UserDTO{}, err
	}
	return users, nil
}

func (u *userService) CreateUser(user model.User) error {
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
