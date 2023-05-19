package s_about

import (
	"strconv"
	"toko-buah/model/m_about"
	"toko-buah/repository/r_about"

	"github.com/gin-gonic/gin"
)

type UpdateAboutService struct {
	updateAboutRepo r_about.UpdateAboutRepository
}

func NewUpdateAboutService(updateAboutRepo r_about.UpdateAboutRepository) *UpdateAboutService {
	return &UpdateAboutService{updateAboutRepo}
}

func (s *UpdateAboutService) UpdateAbout(ctx *gin.Context, about m_about.About) (*m_about.About, error) {
	// Mendapatkan ID about dari parameter route
	idAboutStr := ctx.Param("about_id")

	// Konversi tipe data ID about dari string ke int
	idAbout, err := strconv.Atoi(idAboutStr)
	if err != nil {
		return nil, err
	}

	// Set ID about ke dalam struct about
	about.AboutID = idAbout

	// Memanggil repository untuk update about
	updatedAbout, err := s.updateAboutRepo.UpdateAbout(&about)
	if err != nil {
		return nil, err
	}

	return updatedAbout, nil
}
