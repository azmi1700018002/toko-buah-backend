package r_newarrival

import (
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"
)

type DeleteNewArrivalRepository interface {
	DeleteNewArrivalByID(id int) error
}

type newarrivalDeleteRepository struct{}

func NewDeleteNewArrivalRepository() DeleteNewArrivalRepository {
	return &newarrivalDeleteRepository{}
}

func (*newarrivalDeleteRepository) DeleteNewArrivalByID(id int) (err error) {
	// Menghapus data newarrival dari tabel NewArrival
	if err = db.Server().Unscoped().Where("new_arrival_id = ?", id).Delete(&m_newarrival.NewArrival{}).Error; err != nil {
		return err
	}

	return nil
}
