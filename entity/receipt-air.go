package entity

type ReceiptAirTable interface {
	TableName() string
}

type ReceiptAir struct {
	IDResiUdara uint64 `gorm:"primary_key:auto_increment" json:"id_resi_udara"`
	NomorResi   string `gorm:"type:varchar(255)" json:"nomor_resi"`
	NamaBarang  string `gorm:"type:text" json:"nama_barang"`
	TanggalResi string `gorm:"type:date" json:"tanggal_resi"`
	Ctns        int64  `gorm:"type:int" json:"ctns"`
	Berat       int64  `gorm:"type:int" json:"berat"`
}

func (ReceiptAir) TableName() string {
	return "resi_udara"
}
