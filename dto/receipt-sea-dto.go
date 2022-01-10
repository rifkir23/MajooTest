package dto

/*Handle Json*/
type ReceiptSeaJson struct {
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

/*Receipt Delay*/
type ReceiptDelayListDTO struct {
	IDResi  uint64 `json:"id"`
	Tanggal string `json:"date"`
	Nomor   string `json:"receiptNumber"`
	Status  string `status:"status"`
}

/*Delay list*/
type DelayListDTO struct {
	Total     int64                 `json:"total"`
	Page      int64                 `json:"page"`
	TotalPage int64                 `json:"totalPage"`
	Type      string                `json:"type"`
	Receipt   []ReceiptDelayListDTO `json:"receipt"`
}
