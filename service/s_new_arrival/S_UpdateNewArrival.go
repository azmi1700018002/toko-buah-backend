package s_newarrival

import (
	"mime/multipart"
	"strconv"
	m_newarrival "toko-buah/model/m_new_arrival"
	r_newarrival "toko-buah/repository/r_new_arrival"

	"github.com/gin-gonic/gin"
)

type UpdateNewArrivalService struct {
	updateNewArrivalRepo r_newarrival.UpdateNewArrivalRepository
}

func NewUpdateNewArrivalService(updateNewArrivalRepo r_newarrival.UpdateNewArrivalRepository) *UpdateNewArrivalService {
	return &UpdateNewArrivalService{updateNewArrivalRepo}
}

func (s *UpdateNewArrivalService) UpdateNewArrival(ctx *gin.Context, newarrival m_newarrival.NewArrival, gambarFile *multipart.FileHeader) (*m_newarrival.NewArrival, error) {
	// Mendapatkan ID newarrival dari parameter route
	idNewArrivalStr := ctx.Param("new_arrival_id")

	// Konversi tipe data ID newarrival dari string ke int
	idNewArrival, err := strconv.Atoi(idNewArrivalStr)
	if err != nil {
		return nil, err
	}

	// Set ID newarrival ke dalam struct newarrival
	newarrival.NewArrivalID = idNewArrival

	// Memanggil repository untuk update newarrival
	updatedNewArrival, err := s.updateNewArrivalRepo.UpdateNewArrival(&newarrival, gambarFile)
	if err != nil {
		return nil, err
	}

	return updatedNewArrival, nil
}
