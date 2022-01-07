package entity

type ContainerTable interface {
	TableName() string
}

type Container struct {
	IDContainer        uint64 `gorm:"primary_key:auto_increment" json:"id_container"`
	IDRts              string `json:"id_rts"`
	Nomor              string `json:"nomor"`
	Kode               string `json:"kode"`
	Status             int16  `json:"status"`
	TglLoading         string `json:"tgl_loading"`
	TglClosing         string `json:"tgl_closing"`
	TglEta             string `json:"tgl_eta"`
	TglAntriKapal      string `json:"tgl_antri_kapal"`
	TglAturKapal       string `json:"tgl_atur_kapal"`
	TglEstDumai        string `json:"tgl_est_dumai"`
	TglPib             string `json:"tgl_pib"`
	TglNotul           string `json:"tgl_notul"`
	TanggalBerangkatC  string `json:"tanggal_berangkat_c"`
	TanggalMonitoringC string `json:"tanggal_monitoring_c"`
	TanggalArrivedC    string `json:"tanggal_arrived_c"`
}

func (Container) TableName() string {
	return "container"
}
