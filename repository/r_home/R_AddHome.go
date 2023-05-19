package r_home

import (
	"toko-buah/config/db"
	"toko-buah/model/m_home"
)

type AddHomeRepository interface {
	AddHome(home *m_home.Home) (*m_home.Home, error)
}

type addHomeRepository struct{}

func NewAddHomeRepository() AddHomeRepository {
	return &addHomeRepository{}
}

func (r *addHomeRepository) AddHome(home *m_home.Home) (*m_home.Home, error) {
	err := db.Server().Create(&home).Error
	if err != nil {
		return nil, err
	}
	return home, nil
}
