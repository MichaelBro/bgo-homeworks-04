package card

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
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{
		BankName: bankName,
		Cards:    nil,
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

func (service Service) SearchByNumber(number string) *Card {
	for _, card := range service.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}
