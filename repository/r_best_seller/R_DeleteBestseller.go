package r_bestseller

import (
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_bestseller "toko-buah/model/m_best_seller"
)

type DeleteBestsellerRepository interface {
	DeleteBestsellerByID(id int) error
}

type bestsellerDeleteRepository struct{}

func NewDeleteBestsellerRepository() DeleteBestsellerRepository {
	return &bestsellerDeleteRepository{}
}

func (*bestsellerDeleteRepository) DeleteBestsellerByID(id int) (err error) {
	// Mengambil data bestseller dari database
	bestseller := &m_bestseller.Bestseller{}
	if err = db.Server().First(bestseller, id).Error; err != nil {
		return err
	}

	// Membangun path absolut dari file gambar
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/bestseller/"
	gambarPath := filepath.Join(basePath, bestseller.Gambar)

	// Hapus file gambar
	err = os.Remove(gambarPath)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam menghapus gambar
		return err
	}

	// Menghapus data bestseller dari tabel Bestseller
	if err = db.Server().Unscoped().Delete(bestseller).Error; err != nil {
		return err
	}

	return nil
}
