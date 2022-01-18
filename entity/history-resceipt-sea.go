package entity

type ReceiptSeaHistoryTable interface {
	TableName() string
}

type ReceiptSeaHistory struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ResiId   string `json:"resi_id"`
	Tanggal  int32  `json:"tanggal"`
	TipeResi string `json:"tipe_resi"`
}

func (ReceiptSeaHistory) TableName() string {
	return "history_date_status"
}
