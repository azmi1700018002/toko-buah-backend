package s_produk

import "toko-buah/repository/r_produk"

type DeleteProdukService interface {
	DeleteProdukByID(id int) error
}

type deleteProdukService struct {
	deleteProdukRepo r_produk.DeleteProdukRepository
}

func NewDeleteProdukService(deleteProdukRepo r_produk.DeleteProdukRepository) DeleteProdukService {
	return &deleteProdukService{deleteProdukRepo}
}

func (s *deleteProdukService) DeleteProdukByID(id int) error {
	err := s.deleteProdukRepo.DeleteProdukByID(id)
	if err != nil {
		return err
	}
	return nil
}
