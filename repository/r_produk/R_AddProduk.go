package r_produk

import (
	"mime/multipart"
	"toko-buah/config/db"
	"toko-buah/config/helper"
	"toko-buah/model/m_produk"
)

type AddProdukRepository interface {
	AddProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error)
}

type addProdukRepository struct{}

func NewAddProdukRepository() AddProdukRepository {
	return &addProdukRepository{}
}

func (r *addProdukRepository) AddProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error) {
	// Lakukan penyimpanan file gambar
	// gambarPath, err := helper.SaveFile(gambarFile, "C:/laragon/www/Toko-Buah-Admin/gambar/produk/")
	gambarPath, err := helper.SaveFile(gambarFile, "gambar/produk/")
	if err != nil {
		return nil, err
	}

	produk.Gambar = gambarPath

	// Simpan produk ke database
	err = db.Server().Create(&produk).Error
	if err != nil {
		return nil, err
	}

	return produk, nil
}
