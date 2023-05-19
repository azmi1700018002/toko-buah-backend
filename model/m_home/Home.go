package m_home

type Home struct {
	HomeID    int    `column:"about_id" gorm:"primaryKey;autoIncrement"`
	Subtitle  string `column:"subtitle" validate:"required,min=0"`
	Title     string `column:"judul" validate:"required,min=0"`
	Deskripsi string `column:"deskripsi" validate:"required,min=0"`
}
