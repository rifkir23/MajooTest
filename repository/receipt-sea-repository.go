package repository

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"gorm.io/gorm"
)

//ReceiptSeaRepository is a ....
type ReceiptSeaRepository interface {
	AllReceiptSea() []entity.Resi
	FindReceiptSeaByNumber(resiID string) entity.Resi
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
	db.connection.Where("nomor = ?", resiNumber).First(&receipt_sea)
	return receipt_sea
}

func (db *receiptSeaConnection) AllReceiptSea() []entity.Resi {
	var receipt_sea []entity.Resi
	db.connection.Limit(10).Find(&receipt_sea)
	return receipt_sea
}
