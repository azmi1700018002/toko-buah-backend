package r_produk

import (
	"time"
	"toko-buah/config/db"
	"toko-buah/model/m_produk"
)

type UpdateProdukRepository interface {
	UpdateProduk(produk *m_produk.Produk) (*m_produk.Produk, error)
}

type updateProdukRepository struct{}

func NewUpdateProdukRepository() UpdateProdukRepository {
	return &updateProdukRepository{}
}

func (r *updateProdukRepository) UpdateProduk(produk *m_produk.Produk) (*m_produk.Produk, error) {
	// Check if the user exists
	var existingProduk m_produk.Produk
	if err := db.Server().Where("produk_id = ?", produk.ProdukID).First(&existingProduk).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	produk.CreatedAt = existingProduk.CreatedAt // keep the existing created_at value
	produk.UpdatedAt = time.Now()               // update the updated_at value
	if err := db.Server().Model(&m_produk.Produk{}).
		Where("produk_id = ?", produk.ProdukID).
		Updates(map[string]interface{}{
			"nama":       produk.Nama,
			"deskripsi":  produk.Deskripsi,
			"harga":      produk.Harga,
			"stok":       produk.Stok,
			"created_at": produk.CreatedAt,
			"updated_at": produk.UpdatedAt, // update the updated_at value
		}).Error; err != nil {
		return nil, err
	}

	return produk, nil
}
