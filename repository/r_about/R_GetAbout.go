package r_about

import (
	"toko-buah/config/db"
	"toko-buah/model/m_about"
)

type GetAboutRepository interface {
	GetAllAbout() ([]m_about.About, error)
	GetAboutByID(aboutID int) (*m_about.About, error)
}

type getAboutRepository struct{}

func NewGetAboutRepository() GetAboutRepository {
	return &getAboutRepository{}
}

func (r *getAboutRepository) GetAllAbout() ([]m_about.About, error) {
	var abouts []m_about.About
	result := db.Server().Find(&abouts)
	if result.Error != nil {
		return nil, result.Error
	}
	return abouts, nil
}

func (r *getAboutRepository) GetAboutByID(aboutID int) (*m_about.About, error) {
	var about m_about.About
	result := db.Server().Where("about_id = ?", aboutID).First(&about)
	if result.Error != nil {
		return nil, result.Error
	}
	return &about, nil
}
