package service

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
)

//ReceiptSeaService is a ....
type ReceiptSeaService interface {
	Detail(receiptId int64, containerId int64) dto.ReceiptDetailResult
	Count(customerId int64, cd dto.CountReceiptSea) dto.CountReceiptSea
	List(customerId int64, page int64, limit int64, status string) dto.ReceiptListResult
	ReceiptByContainer(receiptSeaNumber string) []dto.ContainerByReceipt
	Tracking(receiptNumber string, markingCode string) dto.ReceiptDetailResult
}

type receiptSeaService struct {
	receiptSeaRepository repository.ReceiptSeaRepository
}

//NewReceiptSeaService .....
func NewReceiptSeaService(receiptSeaRepo repository.ReceiptSeaRepository) ReceiptSeaService {
	return &receiptSeaService{
		receiptSeaRepository: receiptSeaRepo,
	}
}

func (service *receiptSeaService) Detail(receiptId int64, containerId int64) dto.ReceiptDetailResult {
	return service.receiptSeaRepository.Detail(receiptId, containerId)
}

func (service *receiptSeaService) Count(customerId int64, cd dto.CountReceiptSea) dto.CountReceiptSea {
	return service.receiptSeaRepository.CountReceiptSea(customerId, cd)
}

func (service *receiptSeaService) List(customerId int64, page int64, limit int64, status string) dto.ReceiptListResult {
	return service.receiptSeaRepository.List(customerId, page, limit, status)
}

func (service *receiptSeaService) ReceiptByContainer(receiptSeaNumber string) []dto.ContainerByReceipt {
	return service.receiptSeaRepository.ReceiptByContainer(receiptSeaNumber)
}

func (service *receiptSeaService) Tracking(receiptNumber string, markingCode string) dto.ReceiptDetailResult {
	return service.receiptSeaRepository.Tracking(receiptNumber, markingCode)
}
