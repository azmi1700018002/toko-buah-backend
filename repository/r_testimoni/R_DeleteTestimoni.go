package r_testimoni

import (
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_testimoni "toko-buah/model/m_testimoni"
)

type DeleteTestimoniRepository interface {
	DeleteTestimoniByID(id int) error
}

type testimoniDeleteRepository struct{}

func NewDeleteTestimoniRepository() DeleteTestimoniRepository {
	return &testimoniDeleteRepository{}
}

func (*testimoniDeleteRepository) DeleteTestimoniByID(id int) (err error) {
	// Mengambil data testimoni dari database
	testimoni := &m_testimoni.Testimoni{}
	if err = db.Server().First(testimoni, id).Error; err != nil {
		return err
	}

	// Membangun path absolut dari file gambar
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/testimoni/"
	gambarPath := filepath.Join(basePath, testimoni.Gambar)

	// Hapus file gambar
	err = os.Remove(gambarPath)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam menghapus gambar
		return err
	}

	// Menghapus data testimoni dari tabel Testimoni
	if err = db.Server().Unscoped().Where("testimoni_id = ?", id).Delete(&m_testimoni.Testimoni{}).Error; err != nil {
		return err
	}

	return nil
}
