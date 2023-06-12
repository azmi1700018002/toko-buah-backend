package m_bestseller

type Bestseller struct {
	BestsellerID int    `column:"bestseller_id" gorm:"primaryKey;autoIncrement"`
	Nama         string `column:"nama" validate:"required,min=3,max=100"`
	Deskripsi    string `column:"deskripsi" validate:"max=255"`
	Gambar       string `gorm:"column:gambar"`
}
