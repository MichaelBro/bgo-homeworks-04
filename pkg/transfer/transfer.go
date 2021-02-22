package transfer

import (
	"bgo-homeworks-04/pkg/card"
	"errors"
	"strconv"
	"strings"
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
	ErrInvalidCardNumber    = errors.New("wrong card number")
)

func (service *Service) Card2Card(from, to string, amount int64) (int64, error) {

	if !IsValid(from) || !IsValid(to) {
		return 0, ErrInvalidCardNumber
	}

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

func IsValid(cardNumber string) bool {
	stringSlice := strings.Split(strings.ReplaceAll(cardNumber, " ", ""), "")

	reverseSlice := make([]int, 0)
	for i := len(stringSlice) - 1; i >= 0; i-- { //reverse reading characters from a slice
		number, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			return false
		}
		reverseSlice = append(reverseSlice, number)
	}

	sum := 0
	for i := 0; i < len(reverseSlice); i++ {
		number := reverseSlice[i]
		if (i % 2) != 0 {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
	}
	return (sum % 10) == 0
}
