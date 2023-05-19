package r_newarrival

import (
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"
)

type AddNewArrivalRepository interface {
	AddNewArrival(newarrival *m_newarrival.NewArrival) (*m_newarrival.NewArrival, error)
}

type addNewArrivalRepository struct{}

func NewAddNewArrivalRepository() AddNewArrivalRepository {
	return &addNewArrivalRepository{}
}

func (r *addNewArrivalRepository) AddNewArrival(newarrival *m_newarrival.NewArrival) (*m_newarrival.NewArrival, error) {
	err := db.Server().Create(&newarrival).Error
	if err != nil {
		return nil, err
	}
	return newarrival, nil
}
