package models

type Mahasiswa struct {
	ID        uint `gorm:"primaryKey"`
	Nama      string
	Umur      int
	NIM       string
	TglLahir  string
	Alamat    string
	IDJurusan uint
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
