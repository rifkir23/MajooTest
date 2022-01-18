package entity

type ReceiptSeaTable interface {
	TableName() string
}

type Resi struct {
	IDResi         uint64 `gorm:"primary_key:auto_increment" json:"id_resi"`
	IDResiRts      uint64 `json:"id_resi_rts"`
	Nomor          string `gorm:"type:varchar(255)" json:"nomor"`
	Tanggal        string `gorm:"type:date" json:"tanggal"`
	Supplier       string `gorm:"type:text" json:"supplier"`
	Tel            string `gorm:"type:text" json:"tel"`
	KonfirmasiResi string `gorm:"type:int" json:"konfirmasi_resi"`
	Gudang         int64  `gorm:"type:int" json:"gudang"`

	Giws *[]Giw `gorm:"foreignkey:ResiID;references:IDResi;" json:"giw"`
}

func (Resi) TableName() string {
	return "resi"
}
