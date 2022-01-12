package repository

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"gorm.io/gorm"
)

//ReceiptSeaRepository is a ....
type ReceiptSeaRepository interface {
	AllReceiptSea() []entity.Resi
	FindReceiptSeaByNumber(resiID string) entity.Resi
	CountReceiptSea(cd dto.CountDTO) dto.CountDTO
	Delay(dto.DelayDTO) dto.DelayDTO
	ArrivedSoon(dto.ArrivedSoonDTO) dto.ArrivedSoonDTO
}

type receiptSeaConnection struct {
	connection *gorm.DB
}

//NewReceiptSeaRepository creates an instance ReceiptSeaRepository
func NewReceiptSeaRepository(dbConn *gorm.DB) ReceiptSeaRepository {
	return &receiptSeaConnection{
		connection: dbConn,
	}
}

func (db *receiptSeaConnection) FindReceiptSeaByNumber(resiNumber string) entity.Resi {
	var receipt_sea entity.Resi
	db.connection.Preload("Giws").Preload("Giws.Container").Where("nomor = ?", resiNumber).First(&receipt_sea)
	return receipt_sea
}

func (db *receiptSeaConnection) AllReceiptSea() []entity.Resi {
	var receipt_sea []entity.Resi
	db.connection.Limit(10).Find(&receipt_sea)
	return receipt_sea
}

func (db *receiptSeaConnection) CountReceiptSea(cd dto.CountDTO) dto.CountDTO {
	var giw entity.Giw
	var countDelay int64
	var countArrivedSoon int64
	var countOtw int64

	db.connection.Model(&giw).Joins("LEFT JOIN container on giw.container_id = container.id_rts").Group("resi_id").Where("container.status = ?", "8").Count(&countDelay)
	db.connection.Model(&giw).Joins("LEFT JOIN container on giw.container_id = container.id_rts").Group("resi_id").Where("container.status = ?", "4").Count(&countArrivedSoon)
	db.connection.Model(&giw).Joins("LEFT JOIN container on giw.container_id = container.id_rts").Group("resi_id").Where("container.status = ?", "3").Count(&countOtw)

	cd.Delay = countDelay
	cd.ArrivedSoon = countArrivedSoon
	cd.Otw = countOtw

	return cd
}

func (db *receiptSeaConnection) Delay(dto.DelayDTO) dto.DelayDTO {
	var giw entity.Giw
	var receiptsDelay []dto.ReceiptDelayListDTO
	var countDelay int64
	db.connection.Model(&giw).Raw("SELECT id_resi,resi.tanggal,resi.nomor FROM giw " +
		"LEFT JOIN resi on giw.resi_id = resi.id_resi " +
		"LEFT JOIN container on giw.container_id = container.id_rts " +
		"LEFT JOIN delay on container.id_rts = delay.id_container_rts " +
		"WHERE (container.status = 3 AND delay.tipe = 1) OR (container.status = 8 AND delay.tipe = 2)").Limit(10).Find(&receiptsDelay).Count(&countDelay)
	results := dto.DelayDTO{
		Total:     10,
		Page:      10,
		TotalPage: countDelay,
		Type:      "Delay",
		Receipt:   receiptsDelay,
	}

	return results
}

func (db *receiptSeaConnection) ArrivedSoon(dto.ArrivedSoonDTO) dto.ArrivedSoonDTO {
	var giw entity.Giw
	var receiptsArrivedSoon []dto.ReceiptArrivedSoonDTO
	var countArrivedSoon int64
	db.connection.Model(&giw).Raw("SELECT id_resi,resi.tanggal,resi.nomor FROM giw " +
		"LEFT JOIN resi on giw.resi_id = resi.id_resi " +
		"LEFT JOIN container on giw.container_id = container.id_rts " +
		"LEFT JOIN delay on container.id_rts = delay.id_container_rts " +
		"WHERE container.status = 4 limit 10").Find(&receiptsArrivedSoon).Count(&countArrivedSoon)
	results := dto.ArrivedSoonDTO{
		Total:     10,
		Page:      10,
		TotalPage: countArrivedSoon,
		Type:      "ArrivedSoon",
		Receipt:   receiptsArrivedSoon,
	}

	return results
}
