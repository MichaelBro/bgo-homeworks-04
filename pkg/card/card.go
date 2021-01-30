package card

import (
	"errors"
	"strings"
)

type Card struct {
	Id             int64
	Issuer         string
	FirstNameOwner string
	LastNameOwner  string
	Balance        int64
	Currency       string
	Number         string
	Icon           string
}

type Service struct {
	BankName     string
	NumberPrefix string
	Cards        []*Card
}

func NewService(bankName, numberPrefix string) *Service {
	return &Service{
		BankName:     bankName,
		NumberPrefix: numberPrefix,
		Cards:        nil,
	}
}

func (service *Service) IssueCard(
	id int64,
	issuer string,
	firstNameOwner string,
	lastNameOwner string,
	balance int64,
	currency string,
	number string,
) *Card {
	card := &Card{
		Id:             id,
		Issuer:         issuer,
		FirstNameOwner: firstNameOwner,
		LastNameOwner:  lastNameOwner,
		Balance:        balance,
		Currency:       currency,
		Number:         number,
		Icon:           "icon.png",
	}
	service.Cards = append(service.Cards, card)
	return card
}

var ErrCardNotFound = errors.New("card not found")

func (service *Service) SearchByNumber(number string) (*Card, error) {
	for _, card := range service.Cards {
		if card.Number == number && BelongsToABank(number, service.NumberPrefix) {
			return card, nil
		}
	}
	return nil, ErrCardNotFound
}

func BelongsToABank(numberCard, prefix string) bool {
	return strings.HasPrefix(numberCard, prefix)
}
