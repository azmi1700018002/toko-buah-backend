package s_newarrival

import (
	m_newarrival "toko-buah/model/m_new_arrival"
	r_newarrival "toko-buah/repository/r_new_arrival"

	"github.com/gin-gonic/gin"
)

type AddNewArrivalService struct {
	addNewArrivalRepo r_newarrival.AddNewArrivalRepository
}

func NewAddNewArrivalService(addNewArrivalRepo r_newarrival.AddNewArrivalRepository) *AddNewArrivalService {
	return &AddNewArrivalService{addNewArrivalRepo}
}

func (s *AddNewArrivalService) AddNewArrival(ctx *gin.Context, newarrival m_newarrival.NewArrival) (*m_newarrival.NewArrival, error) {
	updatedProduct, err := s.addNewArrivalRepo.AddNewArrival(&newarrival)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}
