package dto

import "time"

type DelayOtw struct {
	IdDelay        uint64    `json:"id_delay"`
	IdContainerRts string    `json:"id_container_rts"`
	Tipe           int32     `json:"tipe"`
	TglDelay       time.Time `json:"tgl_delay"`
	Keterangan     string    `json:"keterangan"`
}

type DelayEta struct {
	IdDelay        uint64    `json:"id_delay"`
	IdContainerRts string    `json:"id_container_rts"`
	Tipe           int32     `json:"tipe"`
	TglDelay       time.Time `json:"tgl_delay"`
	Keterangan     string    `json:"keterangan"`
}

type DelayOtwLast struct {
	IdDelay        uint64    `json:"id_delay"`
	IdContainerRts string    `json:"id_container_rts"`
	Tipe           int32     `json:"tipe"`
	TglDelay       time.Time `json:"tgl_delay"`
	Keterangan     string    `json:"keterangan"`
}

type DelayEtaLast struct {
	IdDelay        uint64    `json:"id_delay"`
	IdContainerRts string    `json:"id_container_rts"`
	Tipe           int32     `json:"tipe"`
	TglDelay       time.Time `json:"tgl_delay"`
	Keterangan     string    `json:"keterangan"`
}
