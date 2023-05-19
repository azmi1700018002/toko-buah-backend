package s_newarrival

import (
	m_newarrival "toko-buah/model/m_new_arrival"
	r_newarrival "toko-buah/repository/r_new_arrival"
)

type GetNewArrivalService interface {
	GetAllNewArrival() ([]m_newarrival.NewArrival, error)
	GetNewArrivalByID(newarrivalID int) (*m_newarrival.NewArrival, error)
}

type getNewArrivalService struct {
	getNewArrivalRepository r_newarrival.GetNewArrivalRepository
}

func NewGetNewArrivalService(getNewArrivalRepository r_newarrival.GetNewArrivalRepository) GetNewArrivalService {
	return &getNewArrivalService{
		getNewArrivalRepository: getNewArrivalRepository,
	}
}

func (s *getNewArrivalService) GetAllNewArrival() ([]m_newarrival.NewArrival, error) {
	return s.getNewArrivalRepository.GetAllNewArrival()
}

func (s *getNewArrivalService) GetNewArrivalByID(newarrivalID int) (*m_newarrival.NewArrival, error) {
	return s.getNewArrivalRepository.GetNewArrivalByID(newarrivalID)
}
