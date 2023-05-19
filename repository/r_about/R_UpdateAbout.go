package r_about

import (
	"toko-buah/config/db"
	"toko-buah/model/m_about"
)

type UpdateAboutRepository interface {
	UpdateAbout(about *m_about.About) (*m_about.About, error)
}

type updateAboutRepository struct{}

func NewUpdateAboutRepository() UpdateAboutRepository {
	return &updateAboutRepository{}
}

func (r *updateAboutRepository) UpdateAbout(about *m_about.About) (*m_about.About, error) {
	// Check if the user exists
	var existingAbout m_about.About
	if err := db.Server().Where("about_id = ?", about.AboutID).First(&existingAbout).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_about.About{}).
		Where("about_id = ?", about.AboutID).
		Updates(map[string]interface{}{
			"judul":     about.Judul,
			"deskripsi": about.Deskripsi,
		}).Error; err != nil {
		return nil, err
	}
	return about, nil
}
