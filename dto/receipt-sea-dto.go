package dto

/*Handle Json*/
type ReceiptNumber struct {
	ReceiptSeaNumber string `json:"receipt_sea_number"`
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
	Total     int64       `json:"total"`
	Page      int64       `json:"page"`
	TotalPage int64       `json:"totalPage"`
	Type      string      `json:"type"`
	Receipt   interface{} `json:"receipt"`
}

/*Receipt by Container*/
type ContainerByReceiptDTO struct {
	ContainerID int64  `json:"containerID"`
	IDResi      int64  `json:"receiptID"`
	Status      string `json:"status"`
}
