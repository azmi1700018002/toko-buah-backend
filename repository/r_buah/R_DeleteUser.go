package r_buah

import (
	"toko-buah/config/db"
	"toko-buah/model/m_buah"
)

type DeleteBuahRepository interface {
	DeleteBuahByID(id int) error
}

type buahDeleteRepository struct{}

func NewDeleteBuahRepository() DeleteBuahRepository {
	return &buahDeleteRepository{}
}

func (*buahDeleteRepository) DeleteBuahByID(id int) (err error) {
	// Menghapus data buah dari tabel Buah
	if err = db.Server().Unscoped().Where("buah_id = ?", id).Delete(&m_buah.Buah{}).Error; err != nil {
		return err
	}

	return nil
}
