package dto

type ReceiptListByTypeResult struct {
	Pagination interface{} `json:"pagination"`
	Receipt    interface{} `json:"content"`
}

type ReceiptListByType struct {
	ReceiptId     uint64 `json:"id"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	ReceiptNumber string `json:"receiptNumber"`
	Status        string `json:"status"`
}
