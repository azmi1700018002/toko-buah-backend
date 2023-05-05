package s_buah

import (
	"strconv"
	"toko-buah/model/m_buah"
	"toko-buah/repository/r_buah"

	"github.com/gin-gonic/gin"
)

type UpdateBuahService struct {
	updateBuahRepo r_buah.UpdateBuahRepository
}

func NewUpdateBuahService(updateBuahRepo r_buah.UpdateBuahRepository) *UpdateBuahService {
	return &UpdateBuahService{updateBuahRepo}
}

func (s *UpdateBuahService) UpdateBuah(ctx *gin.Context, buah m_buah.Buah) (*m_buah.Buah, error) {
	// Mendapatkan ID buah dari parameter route
	idBuahStr := ctx.Param("buah_id")

	// Konversi tipe data ID buah dari string ke int
	idBuah, err := strconv.Atoi(idBuahStr)
	if err != nil {
		return nil, err
	}

	// Set ID buah ke dalam struct buah
	buah.BuahID = idBuah

	// Memanggil repository untuk update buah
	updatedBuah, err := s.updateBuahRepo.UpdateBuah(&buah)
	if err != nil {
		return nil, err
	}

	return updatedBuah, nil
}
