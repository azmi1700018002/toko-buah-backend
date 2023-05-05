package m_user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID      `gorm:"primary_key; unique; type:uuid; column:id_user; default:uuid_generate_v4()"`
	Username       string         `column:"username"`
	Email          string         `column:"email"`
	Password       string         `column:"password"`
	ProfilePicture string         `column:"profile_picture"`
	CreatedAt      time.Time      `column:"created_at"`
	UpdatedAt      time.Time      `column:"updated_at"`
	DeletedAt      gorm.DeletedAt `column:"deleted_at"`
	LastLogin      *time.Time     `column:"last_login"`
}
