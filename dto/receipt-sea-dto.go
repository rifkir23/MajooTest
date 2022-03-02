package dto

import "time"

/*Handle Json*/
type ReceiptNumber struct {
	ReceiptSeaNumber string `json:"receiptSeaNumber"`
}

/*Receipt Sea Response*/
type ReceiptSeaResponse struct {
	ReceiptSeaNumber string `json:"receipt_sea_number"`
}

/*Count Receipt Sea */
type CountReceiptSea struct {
	Delay       int64 `json:"delay"`
	ArrivedSoon int64 `json:"arrivedSoon"`
	Otw         int64 `json:"onTheWay"`
}

/*Receipt List*/
type ReceiptList struct {
	ReceiptSeaId     uint64 `json:"id"`
	Date             string `json:"date"`
	ReceiptSeaNumber string `json:"receiptNumber"`
	Status           string `json:"status"`
}

/*Receipt list*/
type ReceiptListResult struct {
	Pagination interface{} `json:"pagination"`
	Receipt    interface{} `json:"content"`
}

type Pagination struct {
	TotalElement int64 `json:"totalElement"`
	CurrentPage  int64 `json:"currentPage"`
	NextPage     int64 `json:"nextPage"`
	PrevPage     int64 `json:"prevPage"`
	TotalPage    int64 `json:"totalPage"`
}

/*Receipt by Container*/
type ContainerByReceipt struct {
	ContainerId  int64  `json:"containerID"`
	ReceiptSeaId int64  `json:"receiptID"`
	Status       string `json:"status"`
}

type ReceiptDetailResult struct {
	ReceiptDetail        interface{} `json:"receipt"`
	BarcodeDetailReceipt interface{} `json:"barcode"`
	StatusDetailReceipt  interface{} `json:"statusList"`
}

type ReceiptDetail struct {
	ReceiptSeaId     uint64 `json:"id"`
	ReceiptRtsId     uint64 `json:"receiptIdRts"`
	StatusConfirm    string `json:"status"`
	MarkingCode      string `json:"markingCode"`
	ReceiptSeaNumber string `json:"receiptNumber"`
	Date             string `json:"date"`
	Tel              string `json:"phoneNumber"`
	WhatsappNumber   string `json:"whatsappNumber"`
	Note             string `json:"note"`
	Warehouse        string `json:"warehouseLocation"`
	InsuranceNumber  string `json:"insuranceNumber"`
}

type ReceiptDetailTracking struct {
	ReceiptSeaId     uint64 `json:"id"`
	ReceiptRtsId     uint64 `json:"receiptIdRts"`
	StatusConfirm    string `json:"status"`
	MarkingCode      string `json:"markingCode"`
	ReceiptSeaNumber string `json:"receiptNumber"`
	Date             string `json:"date"`
	Tel              string `json:"phoneNumber"`
	WhatsappNumber   string `json:"whatsappNumber"`
	Note             string `json:"note"`
	Warehouse        string `json:"warehouseLocation"`
	InsuranceNumber  string `json:"insuranceNumber"`
	ContainerId      string `json:"containerId"`
}

type BarcodeDetailReceipt struct {
	TotalCartons string      `json:"totalCartons"`
	TotalQty     string      `json:"totalQty"`
	TotalValue   string      `json:"totalValue"`
	TotalVolume  string      `json:"totalVolume"`
	TotalWeight  string      `json:"totalWeight"`
	BarcodeList  interface{} `json:"barcodeList"`
}

type BarcodeList struct {
	ID        uint64  `json:"id"`
	Nomor     string  `json:"barcode"`
	Qty       string  `json:"qty"`
	Nilai     string  `json:"value"`
	Volume    string  `json:"volume"`
	Berat     string  `json:"weight"`
	HargaJual float64 `json:"sellingPrice"`
}

type StatusDetailReceipt struct {
	//ID           uint64 `json:"id"`
	Date         time.Time `json:"date"`
	ProcessTitle string    `json:"processTitle"`
	Description  string    `json:"description"`
}
