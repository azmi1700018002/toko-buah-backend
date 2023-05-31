package m_newarrival

type NewArrival struct {
	NewArrivalID int     `column:"new_arrival_id" gorm:"primaryKey;autoIncrement"`
	Nama         string  `column:"nama" validate:"required,min=3,max=100"`
	Deskripsi    string  `column:"deskripsi" validate:"max=255"`
	HargaAwal    float64 `column:"harga_awal" validate:"required,min=0"`
	HargaPromo   float64 `column:"harga_promo" validate:"required,min=0"`
	Gambar       string  `gorm:"column:gambar"`
}
