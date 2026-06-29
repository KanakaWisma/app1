package models

type Jurusan struct {
	IDJurusan   uint `gorm:"primaryKey"`
	NamaJurusan string
	Fakultas    string
	Jenjang     string
}

func (Jurusan) TableName() string {
	return "jurusan"
}
