package r_bestseller

import (
	"toko-buah/config/db"
	m_bestseller "toko-buah/model/m_best_seller"
)

type GetBestsellerRepository interface {
	GetAllBestseller() ([]m_bestseller.Bestseller, error)
	GetBestsellerByID(bestsellerID int) (*m_bestseller.Bestseller, error)
}

type getBestsellerRepository struct{}

func NewGetBestsellerRepository() GetBestsellerRepository {
	return &getBestsellerRepository{}
}

func (r *getBestsellerRepository) GetAllBestseller() ([]m_bestseller.Bestseller, error) {
	var bestsellers []m_bestseller.Bestseller
	result := db.Server().Find(&bestsellers)
	if result.Error != nil {
		return nil, result.Error
	}
	return bestsellers, nil
}

func (r *getBestsellerRepository) GetBestsellerByID(bestsellerID int) (*m_bestseller.Bestseller, error) {
	var bestseller m_bestseller.Bestseller
	result := db.Server().Where("bestseller_id = ?", bestsellerID).First(&bestseller)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bestseller, nil
}
