package repository

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserGormSql(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAllUsers() ([]dto.UserDTO, error) {
	var users []dto.UserDTO
	err := u.db.Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) CreateUser(user model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) LoginUser(user model.User) (model.User, error) {
	err := u.db.Model(&model.User{}).Where("email = ? AND password = ?", user.Email, user.Password).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}