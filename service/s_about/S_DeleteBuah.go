package s_about

import "toko-buah/repository/r_about"

type DeleteAboutService interface {
	DeleteAboutByID(id int) error
}

type deleteAboutService struct {
	deleteAboutRepo r_about.DeleteAboutRepository
}

func NewDeleteAboutService(deleteAboutRepo r_about.DeleteAboutRepository) DeleteAboutService {
	return &deleteAboutService{deleteAboutRepo}
}

func (s *deleteAboutService) DeleteAboutByID(id int) error {
	err := s.deleteAboutRepo.DeleteAboutByID(id)
	if err != nil {
		return err
	}
	return nil
}
