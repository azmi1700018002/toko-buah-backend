package r_auth

import (
	"time"
	"toko-buah/config/db"
	"toko-buah/model/m_user"
)

func QAuthUser(email, password string) (user m_user.User, err error) {

	err = db.Server().Where("email = ?", email).First(&user).Error

	return user, err

}

func UpdateLastLogin(user *m_user.User) error {
	now := time.Now()

	user.LastLogin = &now
	return db.Server().Save(user).Error
}
