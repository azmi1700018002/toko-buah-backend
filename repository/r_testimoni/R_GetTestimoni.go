package r_testimoni

import (
	"toko-buah/config/db"
	m_testimoni "toko-buah/model/m_testimoni"
)

type GetTestimoniRepository interface {
	GetAllTestimoni() ([]m_testimoni.Testimoni, error)
	GetTestimoniByID(testimoniID int) (*m_testimoni.Testimoni, error)
}

type getTestimoniRepository struct{}

func NewGetTestimoniRepository() GetTestimoniRepository {
	return &getTestimoniRepository{}
}

func (r *getTestimoniRepository) GetAllTestimoni() ([]m_testimoni.Testimoni, error) {
	var testimonis []m_testimoni.Testimoni
	result := db.Server().Find(&testimonis)
	if result.Error != nil {
		return nil, result.Error
	}
	return testimonis, nil
}

func (r *getTestimoniRepository) GetTestimoniByID(testimoniID int) (*m_testimoni.Testimoni, error) {
	var testimoni m_testimoni.Testimoni
	result := db.Server().Where("testimoni_id = ?", testimoniID).First(&testimoni)
	if result.Error != nil {
		return nil, result.Error
	}
	return &testimoni, nil
}
