package s_bestseller

import (
	m_bestseller "toko-buah/model/m_best_seller"
	r_bestseller "toko-buah/repository/r_best_seller"

	"github.com/gin-gonic/gin"
)

type AddBestsellerService struct {
	addBestsellerRepo r_bestseller.AddBestsellerRepository
}

func NewAddBestsellerService(addBestsellerRepo r_bestseller.AddBestsellerRepository) *AddBestsellerService {
	return &AddBestsellerService{addBestsellerRepo}
}

func (s *AddBestsellerService) AddBestseller(ctx *gin.Context, bestseller m_bestseller.Bestseller) (*m_bestseller.Bestseller, error) {
	gambarFile, err := ctx.FormFile("gambar")
	if err != nil {
		return nil, err
	}

	updatedBestseller, err := s.addBestsellerRepo.AddBestseller(&bestseller, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedBestseller, nil
}
