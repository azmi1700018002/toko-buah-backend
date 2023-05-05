package r_buah

import (
	"toko-buah/config/db"
	"toko-buah/model/m_buah"
)

type GetBuahRepository interface {
	GetAllBuah() ([]m_buah.Buah, error)
	GetBuahByID(buahID int) (*m_buah.Buah, error)
}

type getBuahRepository struct{}

func NewGetBuahRepository() GetBuahRepository {
	return &getBuahRepository{}
}

func (r *getBuahRepository) GetAllBuah() ([]m_buah.Buah, error) {
	var buahs []m_buah.Buah
	result := db.Server().Find(&buahs)
	if result.Error != nil {
		return nil, result.Error
	}
	return buahs, nil
}

func (r *getBuahRepository) GetBuahByID(buahID int) (*m_buah.Buah, error) {
	var buah m_buah.Buah
	result := db.Server().Where("buah_id = ?", buahID).First(&buah)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buah, nil
}
