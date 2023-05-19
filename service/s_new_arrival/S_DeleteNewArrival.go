package s_newarrival

import r_newarrival "toko-buah/repository/r_new_arrival"

type DeleteNewArrivalService interface {
	DeleteNewArrivalByID(id int) error
}

type deleteNewArrivalService struct {
	deleteNewArrivalRepo r_newarrival.DeleteNewArrivalRepository
}

func NewDeleteNewArrivalService(deleteNewArrivalRepo r_newarrival.DeleteNewArrivalRepository) DeleteNewArrivalService {
	return &deleteNewArrivalService{deleteNewArrivalRepo}
}

func (s *deleteNewArrivalService) DeleteNewArrivalByID(id int) error {
	err := s.deleteNewArrivalRepo.DeleteNewArrivalByID(id)
	if err != nil {
		return err
	}
	return nil
}
