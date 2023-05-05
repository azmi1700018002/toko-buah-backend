package s_buah

import (
	"toko-buah/model/m_buah"
	"toko-buah/repository/r_buah"

	"github.com/gin-gonic/gin"
)

type AddBuahService struct {
	addBuahRepo r_buah.AddBuahRepository
}

func NewAddBuahService(addBuahRepo r_buah.AddBuahRepository) *AddBuahService {
	return &AddBuahService{addBuahRepo}
}

func (s *AddBuahService) AddBuah(ctx *gin.Context, buah m_buah.Buah) (*m_buah.Buah, error) {
	updatedUser, err := s.addBuahRepo.AddBuah(&buah)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
