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
