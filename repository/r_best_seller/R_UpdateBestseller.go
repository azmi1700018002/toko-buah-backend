package r_bestseller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_bestseller "toko-buah/model/m_best_seller"

	"github.com/google/uuid"
)

type UpdateBestsellerRepository interface {
	UpdateBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error)
}

type updateBestsellerRepository struct{}

func NewUpdateBestsellerRepository() UpdateBestsellerRepository {
	return &updateBestsellerRepository{}
}

func (r *updateBestsellerRepository) UpdateBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error) {
	// Check if the product exists
	var existingBestseller m_bestseller.Bestseller
	if err := db.Server().Where("bestseller_id = ?", bestseller.BestsellerID).First(&existingBestseller).Error; err != nil {
		return nil, err
	}

	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/bestseller/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingBestseller.Gambar))

	// Jika ada file gambar baru yang dimasukkan, hapus file lama dari local storage
	if gambarFile != nil {
		if existingBestseller.Gambar != "" {
			err := os.Remove(filepath.Join(basePath, existingBestseller.Gambar))
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Jika tidak ada file gambar baru yang dimasukkan, tetapkan kembali nama gambar sebelumnya
		bestseller.Gambar = existingBestseller.Gambar
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
		bestseller.Gambar = newFileName
	}

	// Update product data in the database
	if err := db.Server().Model(&m_bestseller.Bestseller{}).
		Where("bestseller_id = ?", bestseller.BestsellerID).
		Updates(map[string]interface{}{
			"nama":      bestseller.Nama,
			"deskripsi": bestseller.Deskripsi,
			"gambar":    bestseller.Gambar,
		}).Error; err != nil {
		return nil, err
	}

	return bestseller, nil
}
