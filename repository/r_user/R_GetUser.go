package r_user

import (
	"toko-buah/config/db"
	"toko-buah/model/m_user"

	"github.com/google/uuid"
)

type GetUserRepository interface {
	GetAllUser() ([]m_user.User, error)
	GetUserByID(userID uuid.UUID) (*m_user.User, error)
}

type getUserRepository struct{}

func NewGetUserRepository() GetUserRepository {
	return &getUserRepository{}
}

func (r *getUserRepository) GetAllUser() ([]m_user.User, error) {
	var users []m_user.User
	result := db.Server().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *getUserRepository) GetUserByID(userID uuid.UUID) (*m_user.User, error) {
	var user m_user.User
	result := db.Server().Where("id_user = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
