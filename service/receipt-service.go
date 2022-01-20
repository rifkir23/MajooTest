package service

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
)

//ReceiptService is a ....
type ReceiptService interface {
	List(page int64, limit int64, receiptType string) dto.ReceiptListByTypeResult
}

type receiptService struct {
	receiptRepository repository.ReceiptRepository
}

//NewReceiptService .....
func NewReceiptService(receiptRepo repository.ReceiptRepository) ReceiptService {
	return &receiptService{
		receiptRepository: receiptRepo,
	}
}

func (service *receiptService) List(page int64, limit int64, receiptType string) dto.ReceiptListByTypeResult {
	return service.receiptRepository.List(page, limit, receiptType)
}
