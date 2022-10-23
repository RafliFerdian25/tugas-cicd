package repository

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
)

type UserRepository interface {
	GetAllUsers() ([]dto.UserDTO, error)
	CreateUser(user model.User) error
	LoginUser(user model.User) (model.User, error)
}
