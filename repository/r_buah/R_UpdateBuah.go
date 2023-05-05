package r_buah

import (
	"time"
	"toko-buah/config/db"
	"toko-buah/model/m_buah"
)

type UpdateBuahRepository interface {
	UpdateBuah(buah *m_buah.Buah) (*m_buah.Buah, error)
}

type updateBuahRepository struct{}

func NewUpdateBuahRepository() UpdateBuahRepository {
	return &updateBuahRepository{}
}

func (r *updateBuahRepository) UpdateBuah(buah *m_buah.Buah) (*m_buah.Buah, error) {
	// Check if the user exists
	var existingBuah m_buah.Buah
	if err := db.Server().Where("buah_id = ?", buah.BuahID).First(&existingBuah).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	buah.CreatedAt = existingBuah.CreatedAt // keep the existing created_at value
	buah.UpdatedAt = time.Now()             // update the updated_at value
	if err := db.Server().Model(&m_buah.Buah{}).
		Where("buah_id = ?", buah.BuahID).
		Updates(map[string]interface{}{
			"nama":       buah.Nama,
			"deskripsi":  buah.Deskripsi,
			"harga":      buah.Harga,
			"stok":       buah.Stok,
			"created_at": buah.CreatedAt,
			"updated_at": buah.UpdatedAt, // update the updated_at value
		}).Error; err != nil {
		return nil, err
	}

	return buah, nil
}
