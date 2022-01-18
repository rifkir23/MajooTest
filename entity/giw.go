package entity

type GiwTable interface {
	TableName() string
}

type Giw struct {
	ID          uint64  `gorm:"primary_key:auto_increment" json:"id"`
	ResiID      uint    `json:"resi_id"`
	ContainerID uint    `json:"container_id"`
	Nomor       string  `json:"nomor"`
	Barang      string  `json:"barang"`
	Ctns        int32   `json:"ctns"`
	Qty         int32   `json:"qty"`
	Berat       string  `json:"berat"`
	Volume      string  `json:"volume"`
	Nilai       string  `json:"nilai"`
	Note        string  `json:"note"`
	Kurs        float32 `json:"kurs" sql:"type:decimal(13,0);"`
	Remarks     string  `json:"remarks"`
	Harga       float64 `json:"harga" sql:"type:decimal(13,0);"`
	HargaJual   float64 `json:"harga_jual" sql:"type:decimal(13,0);"`

	Container Container `gorm:"foreignkey:IDRts;references:ContainerID;" json:"container"`
}

func (Giw) TableName() string {
	return "giw"
}
