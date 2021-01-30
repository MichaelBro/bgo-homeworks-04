package transfer

import (
	"bgo-homeworks-04/pkg/card"
	"errors"
)

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

var (
	ErrNotErrNotEnoughMoney = errors.New("not enough money")
)

func (service *Service) Card2Card(from, to string, amount int64) (int64, error) {

	cardFrom, errCardFrom := service.CardSvc.SearchByNumber(from)
	cardTo, errCardTo := service.CardSvc.SearchByNumber(to)

	total := getTotalWithCommission(amount, service.Commission, service.MinCommissionAmount)

	if errCardFrom != nil && errCardTo != nil {
		return total, nil
	}

	if errCardFrom != nil && errCardTo == nil {
		cardTo.Balance += amount
		return total, nil
	}

	if cardFrom.Balance < total {
		return total, ErrNotErrNotEnoughMoney
	}

	if errCardFrom == nil && errCardTo != nil {
		cardFrom.Balance -= total
		return total, nil
	}

	cardFrom.Balance -= total
	cardTo.Balance += total

	return total, nil
}

func getTotalWithCommission(amount int64, serviceCommission float64, minCommissionAmount int64) (total int64) {
	commission := int64(float64(amount) * serviceCommission / 100.0)

	if commission < minCommissionAmount {
		commission = minCommissionAmount
	}

	return amount + commission
}
