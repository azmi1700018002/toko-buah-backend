package s_home

import "toko-buah/repository/r_home"

type DeleteHomeService interface {
	DeleteHomeByID(id int) error
}

type deleteHomeService struct {
	deleteHomeRepo r_home.DeleteHomeRepository
}

func NewDeleteHomeService(deleteHomeRepo r_home.DeleteHomeRepository) DeleteHomeService {
	return &deleteHomeService{deleteHomeRepo}
}

func (s *deleteHomeService) DeleteHomeByID(id int) error {
	err := s.deleteHomeRepo.DeleteHomeByID(id)
	if err != nil {
		return err
	}
	return nil
}
