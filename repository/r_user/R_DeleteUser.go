package r_user

import (
	"toko-buah/config/db"
	"toko-buah/model/m_user"

	"github.com/google/uuid"
)

type DeleteUserRepository interface {
	DeleteUserByID(id uuid.UUID) error
}

type userDeleteRepository struct{}

func NewDeleteUserRepository() DeleteUserRepository {
	return &userDeleteRepository{}
}

func (*userDeleteRepository) DeleteUserByID(id uuid.UUID) (err error) {
	// Menghapus data user dari tabel User
	if err = db.Server().Unscoped().Where("id_user = ?", id).Delete(&m_user.User{}).Error; err != nil {
		return err
	}

	return nil
}
