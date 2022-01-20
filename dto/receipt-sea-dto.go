package dto

/*Handle Json*/
type ReceiptNumber struct {
	ReceiptSeaNumber string `json:"receiptSeaNumber"`
}

/*Receipt Sea Response*/
type ReceiptSeaResponse struct {
	ReceiptSeaNumber string `json:"receipt_sea_number"`
}

/*Count Receipt Sea */
type CountDTO struct {
	Delay       int64 `json:"delay"`
	ArrivedSoon int64 `json:"arrivedSoon"`
	Otw         int64 `json:"onTheWay"`
}

/*Receipt List*/
type ReceiptList struct {
	IDResi  uint64 `json:"id"`
	Tanggal string `json:"date"`
	Nomor   string `json:"receiptNumber"`
	Status  string `json:"status"`
}

/*Receipt list*/
type ReceiptListResultDTO struct {
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
type ContainerByReceiptDTO struct {
	ContainerID int64  `json:"containerID"`
	IDResi      int64  `json:"receiptID"`
	Status      string `json:"status"`
}

type ReceiptDetailResult struct {
	ReceiptDetail        interface{} `json:"receipt"`
	BarcodeDetailReceipt interface{} `json:"barcode"`
	StatusDetailReceipt  interface{} `json:"statusList"`
}

type ReceiptDetail struct {
	IDResi          uint64 `json:"id"`
	IDResiRts       uint64 `json:"receiptIdRts"`
	KonfirmasiResi  string `json:"status"`
	MarkingCode     string `json:"markingCode"`
	Nomor           string `json:"receiptNumber"`
	Tanggal         string `json:"date"`
	Tel             string `json:"phoneNumber"`
	WhatsappNumber  string `json:"whatsappNumber"`
	Note            string `json:"note"`
	Gudang          string `json:"warehouseLocation"`
	InsuranceNumber string `json:"insuranceNumber"`
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
	Date         string `json:"date"`
	ProcessTitle string `json:"processTitle"`
	Description  string `json:"description"`
}
