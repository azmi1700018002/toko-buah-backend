package s_buah

import (
	"toko-buah/model/m_buah"
	"toko-buah/repository/r_buah"
)

type GetBuahService interface {
	GetAllBuah() ([]m_buah.Buah, error)
	GetBuahByID(buahID int) (*m_buah.Buah, error)
}

type getBuahService struct {
	getBuahRepository r_buah.GetBuahRepository
}

func NewGetBuahService(getBuahRepository r_buah.GetBuahRepository) GetBuahService {
	return &getBuahService{
		getBuahRepository: getBuahRepository,
	}
}

func (s *getBuahService) GetAllBuah() ([]m_buah.Buah, error) {
	return s.getBuahRepository.GetAllBuah()
}

func (s *getBuahService) GetBuahByID(buahID int) (*m_buah.Buah, error) {
	return s.getBuahRepository.GetBuahByID(buahID)
}
