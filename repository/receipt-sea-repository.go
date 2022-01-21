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
	var container entity.Container
	var historyReceiptSea entity.ReceiptSeaHistory
	var delay []entity.Delay

	var receiptDetail dto.ReceiptDetail
	var barcodeDetail dto.BarcodeDetailReceipt
	var barcodeList []dto.BarcodeList
	var statusDetail []dto.StatusDetailReceipt
	var receiptDetailResult dto.ReceiptDetailResult
	var delayOtw []dto.DelayOtw
	var delayEta []dto.DelayEta
	var delayOtwLast dto.DelayOtwLast

	db.connection.Model(&receipt_sea).Select("resi.id_resi_rts,resi.id_resi,konfirmasi_resi,'123/WC-tes' as MarkingCode,nomor,tanggal,tel,'081312345678' as WhatsappNumber,resi.note,gudang,invoice_asuransi.jumlah_asuransi as InsuranceNumber").
		Joins("LEFT JOIN invoice_asuransi on resi.id_resi = invoice_asuransi.id_resi").First(&receiptDetail, receiptId)
	db.connection.Model(&giw).Where("resi_id = ?", receiptId).Where("container_id = ?", containerId).Find(&barcodeList)
	db.connection.Model(&giw).Select("sum(ctns) as TotalCartons,sum(qty*ctns) as TotalQty,round(sum(nilai*ctns),3) as TotalValue,round(sum(volume*ctns),3) as TotalVolume,round(sum(berat*ctns),3) as TotalWeight").
		Where("resi_id = ?", receiptId).Where("container_id = ?", containerId).Group("resi_id").Find(&barcodeDetail)
	db.connection.Model(&historyReceiptSea).Select("tanggal as Date,status_resi.nama as ProcessTitle,status_resi.keterangan as Description").Joins("LEFT JOIN status_resi on history_date_status.tipe_resi = status_resi.id").Where("resi_id = ?", receiptDetail.IDResiRts).Find(&statusDetail)
	db.connection.Where("id_rts", containerId).First(&container)
	db.connection.Model(&delay).Where("id_container_rts", containerId).Where("tipe", 1).Order("tgl_delay asc").Find(&delayOtw)
	db.connection.Model(&delay).Where("id_container_rts", containerId).Where("tipe", 2).Order("tgl_delay asc").Find(&delayEta)
	db.connection.Model(&delay).Where("id_container_rts", containerId).Where("tipe", 1).Order("tgl_delay desc").First(&delayOtwLast)

	if container.TglLoading != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglLoading,
			ProcessTitle: "Loading Container",
			Description:  "Barang Sedang dimuat kedalam Container di China.",
		})
	}

	if container.TglClosing != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglClosing,
			ProcessTitle: "Closing Container",
			Description:  "Container sudah dapat jadwal jalan kapal.",
		})
	}

	if container.TanggalBerangkatC != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TanggalBerangkatC,
			ProcessTitle: "Container OTW",
			Description:  "Container sudah dapat jadwal jalan kapal.",
		})
	}

	for _, rowDelayO := range delayOtw {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         rowDelayO.TglDelay,
			ProcessTitle: "Container Delay Otw",
			Description:  rowDelayO.Keterangan,
		})
	}

	//tes := "2009-11-10 23:00:00 +0000 UTC m=+0.000000001"
	//if delayOtwLast.TglDelay != "" && time.Parse("2021-01-01", delayOtwLast.TglDelay) <= time.Now() {
	//	statusDetail = append(statusDetail, dto.StatusDetailReceipt{
	//		Date:         delayOtwLast.TglDelay,
	//		ProcessTitle: "Container OTW",
	//		Description:  "Container sudah berangkat dari pelabuhan China.",
	//	})
	//}

	if container.TglEta != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglEta,
			ProcessTitle: "Container ETA",
			Description:  "Container sudah sampai di Malaysia.",
		})
	}

	for _, rowDelayE := range delayEta {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         rowDelayE.TglDelay,
			ProcessTitle: "Container Delay ETA",
			Description:  rowDelayE.Keterangan,
		})
	}

	if container.TglAntriKapal != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglAntriKapal,
			ProcessTitle: "Container Antri Kapal",
			Description:  "Container sedang proses pemesanan jadwal kapal ke Indonesia.",
		})
	}

	if container.TglEstDumai != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglEstDumai,
			ProcessTitle: "Container Estimasi Dumai",
			Description:  "Container sudah sampai di Indonesia dan sedang proses custom clearance",
		})
	}

	if container.TglPib != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglPib,
			ProcessTitle: "Container PIB",
			Description:  "",
		})
	}

	if container.TglNotul != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TglNotul,
			ProcessTitle: "Container Notul",
			Description:  "",
		})
	}

	if container.TanggalMonitoringC != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TanggalMonitoringC,
			ProcessTitle: "Container Monitoring",
			Description:  "Barang sudah Sudah Bisa diMonitoring",
		})
	}

	if container.TanggalArrivedC != "" {
		statusDetail = append(statusDetail, dto.StatusDetailReceipt{
			Date:         container.TanggalArrivedC,
			ProcessTitle: "Container Arrived",
			Description:  "Tiba di Warehouse Jakarta",
		})
	}

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

	db.connection.Model(&giw).Joins("LEFT JOIN container on giw.container_id = container.id_rts").Joins("LEFT JOIN delay on container.id_rts = delay.id_container_rts").Group("resi_id").Where("(container.status = 3 AND delay.tipe = 1) OR (container.status = 8 AND delay.tipe = 2)").Count(&countDelay)
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
