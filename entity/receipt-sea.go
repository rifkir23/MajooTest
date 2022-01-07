package entity

type ReceiptSeaTable interface {
	TableName() string
}

type Resi struct {
	IDResi          uint64 `gorm:"primary_key:auto_increment" json:"id_resi"`
	Nomor           string `gorm:"type:varchar(255)" json:"nomor"`
	Tanggal         string `gorm:"type:date" json:"tanggal"`
	Supplier        string `gorm:"type:text" json:"supplier"`
	Tel             string `gorm:"type:text" json:"tel"`
	Konfirmasi_resi string `gorm:"type:int" json:"konfirmasi_resi"`

	Giws *[]Giw `gorm:"foreignkey:ResiID;references:IDResi;" json:"giw"`
}

func (Resi) TableName() string {
	return "resi"
}
