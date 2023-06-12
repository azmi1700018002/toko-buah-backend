package s_bestseller

import r_bestseller "toko-buah/repository/r_best_seller"

type DeleteBestsellerService interface {
	DeleteBestsellerByID(id int) error
}

type deleteBestsellerService struct {
	deleteBestsellerRepo r_bestseller.DeleteBestsellerRepository
}

func NewDeleteBestsellerService(deleteBestsellerRepo r_bestseller.DeleteBestsellerRepository) DeleteBestsellerService {
	return &deleteBestsellerService{deleteBestsellerRepo}
}

func (s *deleteBestsellerService) DeleteBestsellerByID(id int) error {
	err := s.deleteBestsellerRepo.DeleteBestsellerByID(id)
	if err != nil {
		return err
	}
	return nil
}
