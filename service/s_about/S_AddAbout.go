package s_about

import (
	"toko-buah/model/m_about"
	"toko-buah/repository/r_about"

	"github.com/gin-gonic/gin"
)

type AddAboutService struct {
	addAboutRepo r_about.AddAboutRepository
}

func NewAddAboutService(addAboutRepo r_about.AddAboutRepository) *AddAboutService {
	return &AddAboutService{addAboutRepo}
}

func (s *AddAboutService) AddAbout(ctx *gin.Context, about m_about.About) (*m_about.About, error) {
	updatedUser, err := s.addAboutRepo.AddAbout(&about)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
