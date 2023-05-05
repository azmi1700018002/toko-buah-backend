package s_user

import (
	"toko-buah/model/m_user"
	r_user "toko-buah/repository/r_user"

	"github.com/google/uuid"
)

type GetUserService interface {
	GetAllUser() ([]m_user.User, error)
	GetUserByID(userID uuid.UUID) (*m_user.User, error)
}

type getUserService struct {
	getUserRepository r_user.GetUserRepository
}

func NewGetUserService(getUserRepository r_user.GetUserRepository) GetUserService {
	return &getUserService{
		getUserRepository: getUserRepository,
	}
}

func (s *getUserService) GetAllUser() ([]m_user.User, error) {
	return s.getUserRepository.GetAllUser()
}

func (s *getUserService) GetUserByID(userID uuid.UUID) (*m_user.User, error) {
	return s.getUserRepository.GetUserByID(userID)
}
