package s_testimoni

import (
	m_testimoni "toko-buah/model/m_testimoni"
	r_testimoni "toko-buah/repository/r_testimoni"
)

type GetTestimoniService interface {
	GetAllTestimoni() ([]m_testimoni.Testimoni, error)
	GetTestimoniByID(testimoniID int) (*m_testimoni.Testimoni, error)
}

type getTestimoniService struct {
	getTestimoniRepository r_testimoni.GetTestimoniRepository
}

func NewGetTestimoniService(getTestimoniRepository r_testimoni.GetTestimoniRepository) GetTestimoniService {
	return &getTestimoniService{
		getTestimoniRepository: getTestimoniRepository,
	}
}

func (s *getTestimoniService) GetAllTestimoni() ([]m_testimoni.Testimoni, error) {
	return s.getTestimoniRepository.GetAllTestimoni()
}

func (s *getTestimoniService) GetTestimoniByID(testimoniID int) (*m_testimoni.Testimoni, error) {
	return s.getTestimoniRepository.GetTestimoniByID(testimoniID)
}
