package s_testimoni

import r_testimoni "toko-buah/repository/r_testimoni"

type DeleteTestimoniService interface {
	DeleteTestimoniByID(id int) error
}

type deleteTestimoniService struct {
	deleteTestimoniRepo r_testimoni.DeleteTestimoniRepository
}

func NewDeleteTestimoniService(deleteTestimoniRepo r_testimoni.DeleteTestimoniRepository) DeleteTestimoniService {
	return &deleteTestimoniService{deleteTestimoniRepo}
}

func (s *deleteTestimoniService) DeleteTestimoniByID(id int) error {
	err := s.deleteTestimoniRepo.DeleteTestimoniByID(id)
	if err != nil {
		return err
	}
	return nil
}
