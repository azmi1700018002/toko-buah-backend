package m_about

type About struct {
	AboutID   int    `column:"about_id" gorm:"primaryKey;autoIncrement"`
	Judul     string `column:"judul" validate:"required,min=0"`
	Deskripsi string `column:"deskripsi" validate:"required,min=0"`
}
