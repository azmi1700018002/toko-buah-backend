package r_newarrival

import (
	"os"
	"path/filepath"
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
	// Mengambil data produk dari database
	produk := &m_newarrival.NewArrival{}
	if err = db.Server().First(produk, id).Error; err != nil {
		return err
	}

	// Membangun path absolut dari file gambar
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/newarrival/"
	gambarPath := filepath.Join(basePath, produk.Gambar)

	// Hapus file gambar
	err = os.Remove(gambarPath)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam menghapus gambar
		return err
	}

	// Menghapus data newarrival dari tabel NewArrival
	if err = db.Server().Unscoped().Where("new_arrival_id = ?", id).Delete(&m_newarrival.NewArrival{}).Error; err != nil {
		return err
	}

	return nil
}
