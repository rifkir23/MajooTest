package dto

import "time"

type TransactionReportByOutlet struct {
	MerchantName    string    `json:"merchantName"`
	OutletName      string    `json:"outletName"`
	BillTotal       float64   `json:"omzet"`
	TransactionDate time.Time `json:"transactionDate"`
}

type TransactionReportByMerchant struct {
	MerchantName    string    `json:"merchantName"`
	BillTotal       float64   `json:"omzet"`
	TransactionDate time.Time `json:"transactionDate"`
}
