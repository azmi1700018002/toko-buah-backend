package s_home

import (
	"strconv"
	"toko-buah/model/m_home"
	"toko-buah/repository/r_home"

	"github.com/gin-gonic/gin"
)

type UpdateHomeService struct {
	updateHomeRepo r_home.UpdateHomeRepository
}

func NewUpdateHomeService(updateHomeRepo r_home.UpdateHomeRepository) *UpdateHomeService {
	return &UpdateHomeService{updateHomeRepo}
}

func (s *UpdateHomeService) UpdateHome(ctx *gin.Context, home m_home.Home) (*m_home.Home, error) {
	// Mendapatkan ID home dari parameter route
	idHomeStr := ctx.Param("home_id")

	// Konversi tipe data ID home dari string ke int
	idHome, err := strconv.Atoi(idHomeStr)
	if err != nil {
		return nil, err
	}

	// Set ID home ke dalam struct home
	home.HomeID = idHome

	// Memanggil repository untuk update home
	updatedHome, err := s.updateHomeRepo.UpdateHome(&home)
	if err != nil {
		return nil, err
	}

	return updatedHome, nil
}
