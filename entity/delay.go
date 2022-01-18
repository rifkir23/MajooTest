package entity

type DelayTable interface {
	TableName() string
}

type Delay struct {
	IdDelay        uint64 `gorm:"primary_key:auto_increment" json:"id_delay"`
	IdContainerRts string `json:"id_container_rts"`
	Tipe           int32  `json:"tipe"`
	TglDelay       string `json:"tgl_delay"`
	Keterangan     string `json:"keterangan"`
}

func (Delay) TableName() string {
	return "delay"
}
