package r_produk

import (
	"toko-buah/config/db"
	"toko-buah/model/m_produk"
)

type GetProdukRepository interface {
	GetAllProduk() ([]m_produk.Produk, error)
	GetProdukByID(produkID int) (*m_produk.Produk, error)
}

type getProdukRepository struct{}

func NewGetProdukRepository() GetProdukRepository {
	return &getProdukRepository{}
}

func (r *getProdukRepository) GetAllProduk() ([]m_produk.Produk, error) {
	var produks []m_produk.Produk
	result := db.Server().Find(&produks)
	if result.Error != nil {
		return nil, result.Error
	}
	return produks, nil
}

func (r *getProdukRepository) GetProdukByID(produkID int) (*m_produk.Produk, error) {
	var produk m_produk.Produk
	result := db.Server().Where("produk_id = ?", produkID).First(&produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}
