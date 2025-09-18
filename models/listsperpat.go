package models

type Sparepart struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Jenis  string `json:"jenis"`
	Jumlah int    `json:"jumlah"`
	Harga  int    `json:"harga"`
}

func (Sparepart) TableName() string {
	return "tb_sparepart"
}
