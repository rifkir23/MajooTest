package entity

type StatusContainerTable interface {
	TableName() string
}

type StatusContainer struct {
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	IDRts string `json:"id_rts"`
	Qty   int32  `json:"qty"`
	Berat string `json:"berat"`
}

func (StatusContainer) TableName() string {
	return "status_giw"
}
