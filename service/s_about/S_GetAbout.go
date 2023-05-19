package s_about

import (
	"toko-buah/model/m_about"
	"toko-buah/repository/r_about"
)

type GetAboutService interface {
	GetAllAbout() ([]m_about.About, error)
	GetAboutByID(aboutID int) (*m_about.About, error)
}

type getAboutService struct {
	getAboutRepository r_about.GetAboutRepository
}

func NewGetAboutService(getAboutRepository r_about.GetAboutRepository) GetAboutService {
	return &getAboutService{
		getAboutRepository: getAboutRepository,
	}
}

func (s *getAboutService) GetAllAbout() ([]m_about.About, error) {
	return s.getAboutRepository.GetAllAbout()
}

func (s *getAboutService) GetAboutByID(aboutID int) (*m_about.About, error) {
	return s.getAboutRepository.GetAboutByID(aboutID)
}
