package r_testimoni

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_testimoni "toko-buah/model/m_testimoni"

	"github.com/google/uuid"
)

type UpdateTestimoniRepository interface {
	UpdateTestimoni(testimoni *m_testimoni.Testimoni, gambarFile *multipart.FileHeader) (*m_testimoni.Testimoni, error)
}

type updateTestimoniRepository struct{}

func NewUpdateTestimoniRepository() UpdateTestimoniRepository {
	return &updateTestimoniRepository{}
}

func (r *updateTestimoniRepository) UpdateTestimoni(testimoni *m_testimoni.Testimoni, gambarFile *multipart.FileHeader) (*m_testimoni.Testimoni, error) {
	// Check if the product exists
	var existingTestimoni m_testimoni.Testimoni
	if err := db.Server().Where("testimoni_id = ?", testimoni.TestimoniID).First(&existingTestimoni).Error; err != nil {
		return nil, err
	}

	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/testimoni/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingTestimoni.Gambar))

	// Jika ada file gambar baru yang dimasukkan, hapus file lama dari local storage
	if gambarFile != nil {
		if existingTestimoni.Gambar != "" {
			err := os.Remove(filepath.Join(basePath, existingTestimoni.Gambar))
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Jika tidak ada file gambar baru yang dimasukkan, tetapkan kembali nama gambar sebelumnya
		testimoni.Gambar = existingTestimoni.Gambar
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
		testimoni.Gambar = newFileName
	}

	// Update product data in the database
	if err := db.Server().Model(&m_testimoni.Testimoni{}).
		Where("testimoni_id = ?", testimoni.TestimoniID).
		Updates(map[string]interface{}{
			"nama":      testimoni.Nama,
			"deskripsi": testimoni.Deskripsi,
			"gambar":    testimoni.Gambar,
		}).Error; err != nil {
		return nil, err
	}

	return testimoni, nil
}
