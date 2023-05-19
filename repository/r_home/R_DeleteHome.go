package r_home

import (
	"toko-buah/config/db"
	"toko-buah/model/m_home"
)

type DeleteHomeRepository interface {
	DeleteHomeByID(id int) error
}

type homeDeleteRepository struct{}

func NewDeleteHomeRepository() DeleteHomeRepository {
	return &homeDeleteRepository{}
}

func (*homeDeleteRepository) DeleteHomeByID(id int) (err error) {
	// Menghapus data home dari tabel Home
	if err = db.Server().Unscoped().Where("home_id = ?", id).Delete(&m_home.Home{}).Error; err != nil {
		return err
	}

	return nil
}
