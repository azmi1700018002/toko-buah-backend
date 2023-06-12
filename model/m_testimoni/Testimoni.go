package m_testimoni

type Testimoni struct {
	TestimoniID int    `column:"testimoni_id" gorm:"primaryKey;autoIncrement"`
	Nama        string `column:"nama" validate:"required,min=3,max=100"`
	Deskripsi   string `column:"Deskripsi" validate:"required,min=3,max=100"`
	Gambar      string `gorm:"column:gambar"`
}
