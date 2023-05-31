package r_newarrival

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"

	"github.com/google/uuid"
)

type UpdateNewArrivalRepository interface {
	UpdateNewArrival(newarrival *m_newarrival.NewArrival, gambarFile *multipart.FileHeader) (*m_newarrival.NewArrival, error)
}

type updateNewArrivalRepository struct{}

func NewUpdateNewArrivalRepository() UpdateNewArrivalRepository {
	return &updateNewArrivalRepository{}
}

func (r *updateNewArrivalRepository) UpdateNewArrival(newarrival *m_newarrival.NewArrival, gambarFile *multipart.FileHeader) (*m_newarrival.NewArrival, error) {
	// Check if the user exists
	var existingNewArrival m_newarrival.NewArrival
	if err := db.Server().Where("new_arrival_id = ?", newarrival.NewArrivalID).First(&existingNewArrival).Error; err != nil {
		return nil, err
	}

	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/newarrival/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingNewArrival.Gambar))

	// Jika ada file gambar baru yang dimasukkan, hapus file lama dari local storage
	if gambarFile != nil {
		if existingNewArrival.Gambar != "" {
			err := os.Remove(filepath.Join(basePath, existingNewArrival.Gambar))
			if err != nil {
				return nil, err
			}
		}
	} else {
		// Jika tidak ada file gambar baru yang dimasukkan, tetapkan kembali nama gambar sebelumnya
		newarrival.Gambar = existingNewArrival.Gambar
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

		// Update the arrival's image path
		newarrival.Gambar = newFileName
	}

	// Update user data in the database
	if err := db.Server().Model(&m_newarrival.NewArrival{}).
		Where("new_arrival_id = ?", newarrival.NewArrivalID).
		Updates(map[string]interface{}{
			"nama":        newarrival.Nama,
			"deskripsi":   newarrival.Deskripsi,
			"harga_awal":  newarrival.HargaAwal,
			"harga_promo": newarrival.HargaPromo,
			"gambar":      newarrival.Gambar,
		}).Error; err != nil {
		return nil, err
	}

	return newarrival, nil
}
