package transfer

import "bgo-homeworks-04/pkg/card"

type Service struct {
	CardSvc             *card.Service
	Commission          float64
	MinCommissionAmount int64
}

func NewService(
	cardSvc *card.Service,
	commission float64,
	minCommissionAmount int64,
) *Service {
	return &Service{
		cardSvc,
		commission,
		minCommissionAmount,
	}
}

func (service *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {

	cardFrom := service.CardSvc.SearchByNumber(from)
	cardTo := service.CardSvc.SearchByNumber(to)

	total = getTotalWithCommission(amount, service.Commission, service.MinCommissionAmount)

	ok = false //chek github actions

	if cardFrom != nil {
		if cardFrom.Balance >= total {
			cardFrom.Balance -= total
		} else {
			ok = false
		}
	}

	if cardTo != nil {
		cardTo.Balance += amount
	}

	return
}

func getTotalWithCommission(amount int64, serviceCommission float64, minCommissionAmount int64) (total int64) {
	commission := int64(float64(amount) * serviceCommission / 100.0)

	if commission < minCommissionAmount {
		commission = minCommissionAmount
	}

	return amount + commission
}
