package m_produk

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	ProdukID  int            `column:"produk_id" gorm:"primaryKey;autoIncrement"`
	Nama      string         `column:"nama" validate:"required,min=3,max=100"`
	Deskripsi string         `column:"deskripsi" validate:"max=255"`
	Harga     float64        `column:"harga" validate:"required,min=0"`
	Stok      int            `column:"stok" validate:"required,min=0"`
	CreatedAt time.Time      `column:"created_at"`
	UpdatedAt time.Time      `column:"updated_at"`
	DeletedAt gorm.DeletedAt `column:"deleted_at"`
	Gambar    string         `column:"gambar"`
}
