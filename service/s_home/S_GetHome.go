package s_home

import (
	"toko-buah/model/m_home"
	"toko-buah/repository/r_home"
)

type GetHomeService interface {
	GetAllHome() ([]m_home.Home, error)
	GetHomeByID(homeID int) (*m_home.Home, error)
}

type getHomeService struct {
	getHomeRepository r_home.GetHomeRepository
}

func NewGetHomeService(getHomeRepository r_home.GetHomeRepository) GetHomeService {
	return &getHomeService{
		getHomeRepository: getHomeRepository,
	}
}

func (s *getHomeService) GetAllHome() ([]m_home.Home, error) {
	return s.getHomeRepository.GetAllHome()
}

func (s *getHomeService) GetHomeByID(homeID int) (*m_home.Home, error) {
	return s.getHomeRepository.GetHomeByID(homeID)
}
