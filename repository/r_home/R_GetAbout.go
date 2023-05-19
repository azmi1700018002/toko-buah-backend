package r_home

import (
	"toko-buah/config/db"
	"toko-buah/model/m_home"
)

type GetHomeRepository interface {
	GetAllHome() ([]m_home.Home, error)
	GetHomeByID(homeID int) (*m_home.Home, error)
}

type getHomeRepository struct{}

func NewGetHomeRepository() GetHomeRepository {
	return &getHomeRepository{}
}

func (r *getHomeRepository) GetAllHome() ([]m_home.Home, error) {
	var homes []m_home.Home
	result := db.Server().Find(&homes)
	if result.Error != nil {
		return nil, result.Error
	}
	return homes, nil
}

func (r *getHomeRepository) GetHomeByID(homeID int) (*m_home.Home, error) {
	var home m_home.Home
	result := db.Server().Where("home_id = ?", homeID).First(&home)
	if result.Error != nil {
		return nil, result.Error
	}
	return &home, nil
}
