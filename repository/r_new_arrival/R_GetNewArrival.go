package r_newarrival

import (
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"
)

type GetNewArrivalRepository interface {
	GetAllNewArrival() ([]m_newarrival.NewArrival, error)
	GetNewArrivalByID(newarrivalID int) (*m_newarrival.NewArrival, error)
}

type getNewArrivalRepository struct{}

func NewGetNewArrivalRepository() GetNewArrivalRepository {
	return &getNewArrivalRepository{}
}

func (r *getNewArrivalRepository) GetAllNewArrival() ([]m_newarrival.NewArrival, error) {
	var newarrivals []m_newarrival.NewArrival
	result := db.Server().Find(&newarrivals)
	if result.Error != nil {
		return nil, result.Error
	}
	return newarrivals, nil
}

func (r *getNewArrivalRepository) GetNewArrivalByID(newarrivalID int) (*m_newarrival.NewArrival, error) {
	var newarrival m_newarrival.NewArrival
	result := db.Server().Where("newarrival_id = ?", newarrivalID).First(&newarrival)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newarrival, nil
}
