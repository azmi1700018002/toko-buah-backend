package r_buah

import (
	"toko-buah/config/db"
	"toko-buah/model/m_buah"
)

type AddBuahRepository interface {
	AddBuah(buah *m_buah.Buah) (*m_buah.Buah, error)
}

type addBuahRepository struct{}

func NewAddBuahRepository() AddBuahRepository {
	return &addBuahRepository{}
}

func (r *addBuahRepository) AddBuah(buah *m_buah.Buah) (*m_buah.Buah, error) {
	err := db.Server().Create(&buah).Error
	if err != nil {
		return nil, err
	}
	return buah, nil
}
