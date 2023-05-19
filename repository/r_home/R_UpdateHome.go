package r_home

import (
	"toko-buah/config/db"
	"toko-buah/model/m_home"
)

type UpdateHomeRepository interface {
	UpdateHome(home *m_home.Home) (*m_home.Home, error)
}

type updateHomeRepository struct{}

func NewUpdateHomeRepository() UpdateHomeRepository {
	return &updateHomeRepository{}
}

func (r *updateHomeRepository) UpdateHome(home *m_home.Home) (*m_home.Home, error) {
	// Check if the user exists
	var existingHome m_home.Home
	if err := db.Server().Where("home_id = ?", home.HomeID).First(&existingHome).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_home.Home{}).
		Where("home_id = ?", home.HomeID).
		Updates(map[string]interface{}{
			"subtitle":  home.Subtitle,
			"title":     home.Title,
			"deskripsi": home.Deskripsi,
		}).Error; err != nil {
		return nil, err
	}
	return home, nil
}
