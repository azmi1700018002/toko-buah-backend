package r_produk

import (
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
	// Menghapus data produk dari tabel Produk
	if err = db.Server().Unscoped().Where("produk_id = ?", id).Delete(&m_produk.Produk{}).Error; err != nil {
		return err
	}

	return nil
}
