package r_produk

import (
	"os"
	"path/filepath"
	"toko-buah/config/db"
	"toko-buah/model/m_produk"
)

type DeleteProdukRepository interface {
	DeleteProdukByID(id int) error
}

type produkDeleteRepository struct{}

func NewDeleteProdukRepository() DeleteProdukRepository {
	return &produkDeleteRepository{}
}

func (*produkDeleteRepository) DeleteProdukByID(id int) (err error) {
	// Mengambil data produk dari database
	produk := &m_produk.Produk{}
	if err = db.Server().First(produk, id).Error; err != nil {
		return err
	}

	// Membangun path absolut dari file gambar
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/"
	gambarPath := filepath.Join(basePath, produk.Gambar)

	// Hapus file gambar
	err = os.Remove(gambarPath)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam menghapus gambar
		return err
	}

	// Menghapus data produk dari tabel Produk
	if err = db.Server().Unscoped().Delete(produk).Error; err != nil {
		return err
	}

	return nil
}
