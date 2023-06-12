package s_testimoni

import (
	m_testimoni "toko-buah/model/m_testimoni"
	r_testimoni "toko-buah/repository/r_testimoni"

	"github.com/gin-gonic/gin"
)

type AddTestimoniService struct {
	addTestimoniRepo r_testimoni.AddTestimoniRepository
}

func NewAddTestimoniService(addTestimoniRepo r_testimoni.AddTestimoniRepository) *AddTestimoniService {
	return &AddTestimoniService{addTestimoniRepo}
}

func (s *AddTestimoniService) AddTestimoni(ctx *gin.Context, testimoni m_testimoni.Testimoni) (*m_testimoni.Testimoni, error) {
	gambarFile, err := ctx.FormFile("gambar")
	if err != nil {
		return nil, err
	}

	updatedTestimoni, err := s.addTestimoniRepo.AddTestimoni(&testimoni, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedTestimoni, nil
}
