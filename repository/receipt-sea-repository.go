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
	//Delay(dld dto.DelayListDTO) dto.DelayListDTO
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

//func (db *receiptSeaConnection) Delay(dld dto.DelayListDTO) dto.DelayListDTO {
//	var receipt_sea []entity.Resi
//	db.connection.Limit(10).Find(&receipt_sea)
//	dld.Total = 10
//	dld.Page = 10
//	dld.TotalPage = 100
//	dld.Type = "Delay"
//	dld.Receipt = ""
//	return dld
//}
