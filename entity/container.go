package entity

import "time"

type ContainerTable interface {
	TableName() string
}

type Container struct {
	IDContainer        uint64    `gorm:"primary_key:auto_increment" json:"id_container"`
	IDRts              string    `json:"id_rts"`
	Nomor              string    `json:"nomor"`
	Kode               string    `json:"kode"`
	Status             int16     `json:"status"`
	TglLoading         time.Time `json:"tgl_loading"`
	TglClosing         time.Time `json:"tgl_closing"`
	TglEta             time.Time `json:"tgl_eta"`
	TglAntriKapal      time.Time `json:"tgl_antri_kapal"`
	TglAturKapal       time.Time `json:"tgl_atur_kapal"`
	TglEstDumai        time.Time `json:"tgl_est_dumai"`
	TglPib             time.Time `json:"tgl_pib"`
	TglNotul           time.Time `json:"tgl_notul"`
	TanggalBerangkatC  time.Time `json:"tanggal_berangkat_c"`
	TanggalMonitoringC time.Time `json:"tanggal_monitoring_c"`
	TanggalArrivedC    time.Time `json:"tanggal_arrived_c"`

	StatusContainer StatusContainer `gorm:"foreignkey:ID;references:Status;" json:"statusContainer"`
}

func (Container) TableName() string {
	return "container"
}
