package r_user

import (
	"toko-buah/config/db"
	"toko-buah/config/helper"
	"toko-buah/model/m_user"

	"github.com/google/uuid"
)

type RegisterUserRepository interface {
	RegisterUser(user *m_user.User) error
}

type registerUserRepository struct{}

func NewUserRepository() RegisterUserRepository {
	return &registerUserRepository{}
}

func (r *registerUserRepository) RegisterUser(user *m_user.User) error {
	id_user, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user.UserID = id_user

	// Hash password before save
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Menyimpan user ke dalam database
	if err := db.Server().Create(user).Error; err != nil {
		return err
	}

	return nil
}
