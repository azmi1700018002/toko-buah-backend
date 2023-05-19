package r_about

import (
	"toko-buah/config/db"
	"toko-buah/model/m_about"
)

type AddAboutRepository interface {
	AddAbout(about *m_about.About) (*m_about.About, error)
}

type addAboutRepository struct{}

func NewAddAboutRepository() AddAboutRepository {
	return &addAboutRepository{}
}

func (r *addAboutRepository) AddAbout(about *m_about.About) (*m_about.About, error) {
	err := db.Server().Create(&about).Error
	if err != nil {
		return nil, err
	}
	return about, nil
}
