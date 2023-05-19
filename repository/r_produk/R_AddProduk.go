package r_produk

import (
	"toko-buah/config/db"
	"toko-buah/model/m_produk"
)

type AddProdukRepository interface {
	AddProduk(produk *m_produk.Produk) (*m_produk.Produk, error)
}

type addProdukRepository struct{}

func NewAddProdukRepository() AddProdukRepository {
	return &addProdukRepository{}
}

func (r *addProdukRepository) AddProduk(produk *m_produk.Produk) (*m_produk.Produk, error) {
	err := db.Server().Create(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}
