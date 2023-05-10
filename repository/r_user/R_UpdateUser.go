package r_user

import (
	"time"
	"toko-buah/config/db"
	"toko-buah/config/helper"
	"toko-buah/model/m_user"
)

type UpdateUserRepository interface {
	UpdateUser(user *m_user.User) (*m_user.User, error)
}

type updateUserRepository struct{}

func NewUpdateUserRepository() UpdateUserRepository {
	return &updateUserRepository{}
}

func (r *updateUserRepository) UpdateUser(user *m_user.User) (*m_user.User, error) {
	// Hash password before update
	if user.Password != "" {
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	// Check if the user exists
	var existingUser m_user.User
	if err := db.Server().Where("id_user = ?", user.UserID).First(&existingUser).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	user.CreatedAt = existingUser.CreatedAt // keep the existing created_at value
	user.UpdatedAt = time.Now()             // update the updated_at value
	if err := db.Server().Model(&m_user.User{}).
		Where("id_user = ?", user.UserID).
		Updates(map[string]interface{}{
			"username":        user.Username,
			"email":           user.Email,
			"password":        user.Password,
			"profile_picture": user.ProfilePicture,
			"created_at":      user.CreatedAt,
			"updated_at":      user.UpdatedAt, // update the updated_at value
		}).Error; err != nil {
		return nil, err
	}

	return user, nil
}
