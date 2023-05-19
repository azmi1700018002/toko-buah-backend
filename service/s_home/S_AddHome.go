package s_home

import (
	"toko-buah/model/m_home"
	"toko-buah/repository/r_home"

	"github.com/gin-gonic/gin"
)

type AddHomeService struct {
	addHomeRepo r_home.AddHomeRepository
}

func NewAddHomeService(addHomeRepo r_home.AddHomeRepository) *AddHomeService {
	return &AddHomeService{addHomeRepo}
}

func (s *AddHomeService) AddHome(ctx *gin.Context, home m_home.Home) (*m_home.Home, error) {
	updatedUser, err := s.addHomeRepo.AddHome(&home)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
