package dto

type BodyListReceipt struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Status string `json:"status"`
}
