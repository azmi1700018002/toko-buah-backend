package r_about

import (
	"toko-buah/config/db"
	"toko-buah/model/m_about"
)

type DeleteAboutRepository interface {
	DeleteAboutByID(id int) error
}

type aboutDeleteRepository struct{}

func NewDeleteAboutRepository() DeleteAboutRepository {
	return &aboutDeleteRepository{}
}

func (*aboutDeleteRepository) DeleteAboutByID(id int) (err error) {
	// Menghapus data about dari tabel About
	if err = db.Server().Unscoped().Where("about_id = ?", id).Delete(&m_about.About{}).Error; err != nil {
		return err
	}

	return nil
}
