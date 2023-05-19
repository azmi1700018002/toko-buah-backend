package s_produk

import (
	"toko-buah/model/m_produk"
	"toko-buah/repository/r_produk"
)

type GetProdukService interface {
	GetAllProduk() ([]m_produk.Produk, error)
	GetProdukByID(produkID int) (*m_produk.Produk, error)
}

type getProdukService struct {
	getProdukRepository r_produk.GetProdukRepository
}

func NewGetProdukService(getProdukRepository r_produk.GetProdukRepository) GetProdukService {
	return &getProdukService{
		getProdukRepository: getProdukRepository,
	}
}

func (s *getProdukService) GetAllProduk() ([]m_produk.Produk, error) {
	return s.getProdukRepository.GetAllProduk()
}

func (s *getProdukService) GetProdukByID(produkID int) (*m_produk.Produk, error) {
	return s.getProdukRepository.GetProdukByID(produkID)
}
