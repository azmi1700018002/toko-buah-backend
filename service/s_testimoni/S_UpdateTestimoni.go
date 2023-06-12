package s_testimoni

import (
	"mime/multipart"
	"strconv"
	m_testimoni "toko-buah/model/m_testimoni"
	r_testimoni "toko-buah/repository/r_testimoni"

	"github.com/gin-gonic/gin"
)

type UpdateTestimoniService struct {
	updateTestimoniRepo r_testimoni.UpdateTestimoniRepository
}

func NewUpdateTestimoniService(updateTestimoniRepo r_testimoni.UpdateTestimoniRepository) *UpdateTestimoniService {
	return &UpdateTestimoniService{updateTestimoniRepo}
}

func (s *UpdateTestimoniService) UpdateTestimoni(ctx *gin.Context, testimoni m_testimoni.Testimoni, gambarFile *multipart.FileHeader) (*m_testimoni.Testimoni, error) {
	// Mendapatkan ID testimoni dari parameter route
	idTestimoniStr := ctx.Param("testimoni_id")

	// Konversi tipe data ID testimoni dari string ke int
	idTestimoni, err := strconv.Atoi(idTestimoniStr)
	if err != nil {
		return nil, err
	}

	// Set ID testimoni ke dalam struct testimoni
	testimoni.TestimoniID = idTestimoni

	// Memanggil repository untuk update testimoni
	updatedTestimoni, err := s.updateTestimoniRepo.UpdateTestimoni(&testimoni, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedTestimoni, nil
}
