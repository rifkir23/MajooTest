package entity

type StatusContainerTable interface {
	TableName() string
}

type StatusContainer struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Urutan int64  `json:"urutan"`
}

func (StatusContainer) TableName() string {
	return "status_giw"
}
