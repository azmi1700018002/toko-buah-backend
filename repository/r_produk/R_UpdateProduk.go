package r_produk

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"toko-buah/config/db"
	"toko-buah/model/m_produk"

	"github.com/google/uuid"
)

type UpdateProdukRepository interface {
	UpdateProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error)
}

type updateProdukRepository struct{}

func NewUpdateProdukRepository() UpdateProdukRepository {
	return &updateProdukRepository{}
}

func (r *updateProdukRepository) UpdateProduk(produk *m_produk.Produk, gambarFile *multipart.FileHeader) (*m_produk.Produk, error) {
	// Check if the product exists
	var existingProduk m_produk.Produk
	if err := db.Server().Where("produk_id = ?", produk.ProdukID).First(&existingProduk).Error; err != nil {
		return nil, err
	}

	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingProduk.Gambar))

	// Jika ada file gambar baru yang dimasukkan, hapus file lama dari local storage
	if gambarFile != nil {
		if existingProduk.Gambar != "" {
			err := os.Remove(filepath.Join(basePath, existingProduk.Gambar))
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Jika tidak ada file gambar baru yang dimasukkan, tetapkan kembali nama gambar sebelumnya
		produk.Gambar = existingProduk.Gambar
	}

	// Jika ada file gambar baru yang dimasukkan, proses gambar baru
	if gambarFile != nil {
		// Generate unique filename using UUID for the new file
		fileExt := filepath.Ext(gambarFile.Filename)
		newFileName := uuid.New().String() + fileExt

		// Open the source file
		src, err := gambarFile.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Create the destination file
		savePath := filepath.Join(basePath, newFileName)

		dst, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copy the contents from source to destination
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, err
		}

		// Update the product's image path
		produk.Gambar = newFileName
	}

	produk.CreatedAt = existingProduk.CreatedAt // keep the existing created_at value
	produk.UpdatedAt = time.Now()               // update the updated_at value
	// Update product data in the database
	if err := db.Server().Model(&m_produk.Produk{}).
		Where("produk_id = ?", produk.ProdukID).
		Updates(map[string]interface{}{
			"nama":       produk.Nama,
			"deskripsi":  produk.Deskripsi,
			"harga":      produk.Harga,
			"stok":       produk.Stok,
			"gambar":     produk.Gambar,
			"created_at": produk.CreatedAt,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return nil, err
	}

	return produk, nil
}
