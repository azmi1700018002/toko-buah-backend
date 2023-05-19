package r_newarrival

import (
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"
)

type UpdateNewArrivalRepository interface {
	UpdateNewArrival(newarrival *m_newarrival.NewArrival) (*m_newarrival.NewArrival, error)
}

type updateNewArrivalRepository struct{}

func NewUpdateNewArrivalRepository() UpdateNewArrivalRepository {
	return &updateNewArrivalRepository{}
}

func (r *updateNewArrivalRepository) UpdateNewArrival(newarrival *m_newarrival.NewArrival) (*m_newarrival.NewArrival, error) {
	// Check if the user exists
	var existingNewArrival m_newarrival.NewArrival
	if err := db.Server().Where("new_arrival_id = ?", newarrival.NewArrivalID).First(&existingNewArrival).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_newarrival.NewArrival{}).
		Where("new_arrival_id = ?", newarrival.NewArrivalID).
		Updates(map[string]interface{}{
			"nama":        newarrival.Nama,
			"deskripsi":   newarrival.Deskripsi,
			"harga_awal":  newarrival.HargaAwal,
			"harga_promo": newarrival.HargaPromo,
		}).Error; err != nil {
		return nil, err
	}

	return newarrival, nil
}
