package s_buah

import "toko-buah/repository/r_buah"

type DeleteBuahService interface {
	DeleteBuahByID(id int) error
}

type deleteBuahService struct {
	deleteBuahRepo r_buah.DeleteBuahRepository
}

func NewDeleteBuahService(deleteBuahRepo r_buah.DeleteBuahRepository) DeleteBuahService {
	return &deleteBuahService{deleteBuahRepo}
}

func (s *deleteBuahService) DeleteBuahByID(id int) error {
	err := s.deleteBuahRepo.DeleteBuahByID(id)
	if err != nil {
		return err
	}
	return nil
}
