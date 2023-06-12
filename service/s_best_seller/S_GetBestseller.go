package s_bestseller

import (
	m_bestseller "toko-buah/model/m_best_seller"
	r_bestseller "toko-buah/repository/r_best_seller"
)

type GetBestsellerService interface {
	GetAllBestseller() ([]m_bestseller.Bestseller, error)
	GetBestsellerByID(bestsellerID int) (*m_bestseller.Bestseller, error)
}

type getBestsellerService struct {
	getBestsellerRepository r_bestseller.GetBestsellerRepository
}

func NewGetBestsellerService(getBestsellerRepository r_bestseller.GetBestsellerRepository) GetBestsellerService {
	return &getBestsellerService{
		getBestsellerRepository: getBestsellerRepository,
	}
}

func (s *getBestsellerService) GetAllBestseller() ([]m_bestseller.Bestseller, error) {
	return s.getBestsellerRepository.GetAllBestseller()
}

func (s *getBestsellerService) GetBestsellerByID(bestsellerID int) (*m_bestseller.Bestseller, error) {
	return s.getBestsellerRepository.GetBestsellerByID(bestsellerID)
}
