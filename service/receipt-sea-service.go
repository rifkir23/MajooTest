package service

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
)

//ReceiptSeaService is a ....
type ReceiptSeaService interface {
	All() []entity.Resi
	FindByNumber(resiNumber string) entity.Resi
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

func (service *receiptSeaService) All() []entity.Resi {
	return service.receiptSeaRepository.AllReceiptSea()
}

func (service *receiptSeaService) FindByNumber(resiNumber string) entity.Resi {
	return service.receiptSeaRepository.FindReceiptSeaByNumber(resiNumber)
}
