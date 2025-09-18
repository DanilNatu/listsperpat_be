package models

type Sperpat struct {
    ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Jenis  string `json:"jenis"`
    Jumlah int    `json:"jumlah"`
    Harga  string `json:"harga"`
}

func (Sperpat) TableName() string {
    return "tb_sperpat"
}