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
	Detail(receiptId int64, containerId int64) dto.ReceiptDetailResult
	CountReceiptSea(cd dto.CountDTO) dto.CountDTO
	List(page int64, limit int64, status string) dto.ReceiptListResultDTO
	ReceiptByContainer(resiNumber string) []dto.ContainerByReceiptDTO
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

func (db *receiptSeaConnection) Detail(receiptId int64, containerId int64) dto.ReceiptDetailResult {
	var receipt_sea entity.Resi
	var giw entity.Giw
	var historyReceiptSea entity.ReceiptSeaHistory

	var receiptDetail dto.ReceiptDetail
	var barcodeDetail dto.BarcodeDetailReceipt
	var barcodeList []dto.BarcodeList
	var statusDetail []dto.StatusDetailReceipt
	var receiptDetailResult dto.ReceiptDetailResult

	db.connection.Model(&receipt_sea).Select("resi.id_resi_rts,resi.id_resi,konfirmasi_resi,'123/WC-tes' as MarkingCode,nomor,tanggal,tel,'081312345678' as WhatsappNumber,resi.note,gudang,invoice_asuransi.jumlah_asuransi as InsuranceNumber").
		Joins("LEFT JOIN invoice_asuransi on resi.id_resi = invoice_asuransi.id_resi").First(&receiptDetail, receiptId)
	db.connection.Model(&giw).Where("resi_id = ?", receiptId).Where("container_id = ?", containerId).Find(&barcodeList)
	db.connection.Model(&giw).Select("sum(ctns) as TotalCartons,sum(qty*ctns) as TotalQty,round(sum(nilai*ctns),3) as TotalValue,round(sum(volume*ctns),3) as TotalVolume,round(sum(berat*ctns),3) as TotalWeight").
		Where("resi_id = ?", receiptId).Where("container_id = ?", containerId).Group("resi_id").Find(&barcodeDetail)
	db.connection.Model(&historyReceiptSea).Select("tanggal as Date,status_resi.nama as ProcessTitle,status_resi.keterangan as Description").Joins("LEFT JOIN status_resi on history_date_status.tipe_resi = status_resi.id").Where("resi_id = ?", receiptDetail.IDResiRts).Find(&statusDetail)

	receiptDetailResult.ReceiptDetail = receiptDetail
	barcodeDetail.BarcodeList = barcodeList
	receiptDetailResult.BarcodeDetailReceipt = barcodeDetail
	receiptDetailResult.StatusDetailReceipt = statusDetail

	return receiptDetailResult
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
	var pagination dto.Pagination
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

	/*Current Page*/
	if page >= (countList / limit) {
		pagination.CurrentPage = countList / limit
	} else {
		pagination.CurrentPage = page
	}
	/*Prev Page*/
	if (page-1) > 0 && (countList/limit) > 0 {
		if page >= (countList / limit) {
			pagination.PrevPage = pagination.CurrentPage - 1
		} else {
			pagination.PrevPage = page - 1
		}
	}
	/*Next Page*/
	if (page + 1) < (countList / limit) {
		pagination.NextPage = page + 1
	} else {
		pagination.NextPage = 0
	}
	pagination.TotalElement = countList
	pagination.TotalPage = countList / limit

	results := dto.ReceiptListResultDTO{
		Pagination: pagination,
		Receipt:    receiptList,
	}

	return results
}

func (db *receiptSeaConnection) ReceiptByContainer(resiNumber string) []dto.ContainerByReceiptDTO {
	var giw entity.Giw
	var receipt_by_container_list []dto.ContainerByReceiptDTO
	db.connection.Model(&giw).
		Select("giw.container_id,resi.id_resi,status_giw.status").
		Joins("LEFT JOIN resi on giw.resi_id = resi.id_resi ").
		Joins("LEFT JOIN container on giw.container_id = container.id_rts ").
		Joins("LEFT JOIN status_giw on container.status = status_giw.id ").
		Where("resi.nomor = ?", resiNumber).Group("giw.container_id").Find(&receipt_by_container_list)

	return receipt_by_container_list
}
