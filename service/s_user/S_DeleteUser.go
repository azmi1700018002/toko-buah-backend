package s_user

import (
	r_user "toko-buah/repository/r_user"

	"github.com/google/uuid"
)

type DeleteUserService interface {
	DeleteUserByID(id uuid.UUID) error
}

type deleteUserService struct {
	deleteUserRepo r_user.DeleteUserRepository
}

func NewDeleteUserService(deleteUserRepo r_user.DeleteUserRepository) DeleteUserService {
	return &deleteUserService{deleteUserRepo}
}

func (s *deleteUserService) DeleteUserByID(id uuid.UUID) error {
	err := s.deleteUserRepo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
