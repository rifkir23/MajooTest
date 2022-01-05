package entity

type Tabler interface {
	TableName() string
}

type Resi struct {
	//gorm.Model
	IDResi          uint64 `gorm:"primary_key:auto_increment" json:"id_resi"`
	Nomor           string `gorm:"type:varchar(255)" json:"nomor"`
	Tanggal         string `gorm:"type:date" json:"tanggal"`
	Supplier        string `gorm:"type:text" json:"supplier"`
	Tel             string `gorm:"type:text" json:"tel"`
	Konfirmasi_resi string `gorm:"type:int" json:"konfirmasi_resi"`
	//CustID          uint64 `gorm:"not null" json:"-"`
	//Customer        Customer   `gorm:"foreignkey:CustID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`
}

func (Resi) TableName() string {
	return "resi"
}
