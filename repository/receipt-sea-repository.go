package repository

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"github.com/wilopo-cargo/microservice-receipt-sea/helper"
	"gorm.io/gorm"
)

//ReceiptSeaRepository is a ....
type ReceiptSeaRepository interface {
	AllReceiptSea() []entity.Resi
	FindReceiptSeaByNumber(resiID string) entity.Resi
	CountReceiptSea(cd dto.CountDTO) dto.CountDTO
	List(page int64, limit int64, status string) dto.ReceiptListResultDTO
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

func (db *receiptSeaConnection) List(page int64, limit int64, status string) dto.ReceiptListResultDTO {
	var giw entity.Giw
	var receiptList []dto.ReceiptList
	var countList int64

	if status == "arrivedSoon" {
		db.connection.Model(&giw).Select("id_resi,resi.tanggal,resi.nomor,'"+status+"'as status").
			Joins("LEFT JOIN resi on giw.resi_id = resi.id_resi ").
			Joins("LEFT JOIN container on giw.container_id = container.id_rts ").
			Where("container.status = ?", 4).
			Count(&countList).Scopes(helper.PaginateReceipt(page, limit)).Find(&receiptList)
	} else if status == "delay" {
		db.connection.Model(&giw).Select("id_resi,resi.tanggal,resi.nomor,'" + status + "'as status").
			Joins("LEFT JOIN resi on giw.resi_id = resi.id_resi ").
			Joins("LEFT JOIN container on giw.container_id = container.id_rts ").
			Joins("LEFT JOIN delay on container.id_rts = delay.id_container_rts").
			Where("(container.status = 3 AND delay.tipe = 1) OR (container.status = 8 AND delay.tipe = 2)").
			Count(&countList).Scopes(helper.PaginateReceipt(page, limit)).Find(&receiptList)
	}

	results := dto.ReceiptListResultDTO{
		Total:     countList,
		Page:      page,
		TotalPage: countList / limit,
		Type:      status,
		Receipt:   receiptList,
	}

	return results
}
