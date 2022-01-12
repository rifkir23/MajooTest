package service

import (
	"github.com/wilopo-cargo/microservice-receipt-sea/dto"
	"github.com/wilopo-cargo/microservice-receipt-sea/entity"
	"github.com/wilopo-cargo/microservice-receipt-sea/repository"
)

//ReceiptSeaService is a ....
type ReceiptSeaService interface {
	All() []entity.Resi
	FindByNumber(resiNumber string) entity.Resi
	Count(cd dto.CountDTO) dto.CountDTO
	Delay(dd dto.DelayDTO) dto.DelayDTO
	ArrivedSoon(asd dto.ArrivedSoonDTO) dto.ArrivedSoonDTO
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

func (service *receiptSeaService) Count(cd dto.CountDTO) dto.CountDTO {
	return service.receiptSeaRepository.CountReceiptSea(cd)
}

func (service *receiptSeaService) Delay(dd dto.DelayDTO) dto.DelayDTO {
	return service.receiptSeaRepository.Delay(dd)
}

func (service *receiptSeaService) ArrivedSoon(asd dto.ArrivedSoonDTO) dto.ArrivedSoonDTO {
	return service.receiptSeaRepository.ArrivedSoon(asd)
}
