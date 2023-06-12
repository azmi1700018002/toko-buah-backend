package s_bestseller

import (
	"mime/multipart"
	"strconv"
	m_bestseller "toko-buah/model/m_best_seller"
	r_bestseller "toko-buah/repository/r_best_seller"

	"github.com/gin-gonic/gin"
)

type UpdateBestsellerService struct {
	updateBestsellerRepo r_bestseller.UpdateBestsellerRepository
}

func NewUpdateBestsellerService(updateBestsellerRepo r_bestseller.UpdateBestsellerRepository) *UpdateBestsellerService {
	return &UpdateBestsellerService{updateBestsellerRepo}
}

func (s *UpdateBestsellerService) UpdateBestseller(ctx *gin.Context, bestseller m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error) {
	// Mendapatkan ID bestseller dari parameter route
	idBestsellerStr := ctx.Param("bestseller_id")

	// Konversi tipe data ID bestseller dari string ke int
	idBestseller, err := strconv.Atoi(idBestsellerStr)
	if err != nil {
		return nil, err
	}

	// Set ID bestseller ke dalam struct bestseller
	bestseller.BestsellerID = idBestseller

	// Memanggil repository untuk update bestseller
	updatedBestseller, err := s.updateBestsellerRepo.UpdateBestseller(&bestseller, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedBestseller, nil
}
