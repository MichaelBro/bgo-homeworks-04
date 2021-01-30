package transfer

import (
	"bgo-homeworks-04/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc             *card.Service
		Commission          float64
		MinCommissionAmount int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantErr   error
	}{
		{
			name: "№1 Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             1,
							Issuer:         "Visa",
							FirstNameOwner: "Michael",
							LastNameOwner:  "Bro",
							Balance:        100_000_00,
							Currency:       "RUB",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
						{
							Id:             1,
							Issuer:         "Visa",
							FirstNameOwner: "Alex",
							LastNameOwner:  "Storm",
							Balance:        50_000_00,
							Currency:       "RUB",
							Number:         "5106 2146 1200 5008",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5106 2146 1200 5461",
				to:     "5106 2112 1234 5461",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantErr:   nil,
		},
		{
			name: "№2 Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Visa",
							FirstNameOwner: "Michael",
							LastNameOwner:  "Bro",
							Balance:        5000_00,
							Currency:       "RUB",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
						{
							Id:             0,
							Issuer:         "Visa",
							FirstNameOwner: "Alex",
							LastNameOwner:  "Storm",
							Balance:        50_000_00,
							Currency:       "RUB",
							Number:         "5106 2146 1200 5008",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5106 2146 1200 5461",
				to:     "5106 2146 1200 5008",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantErr:   ErrNotErrNotEnoughMoney,
		},
		{
			name: "№3 Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Maestro",
							FirstNameOwner: "Bob",
							LastNameOwner:  "Down",
							Balance:        16_125_99,
							Currency:       "USD",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5106 2146 1200 5461",
				to:     "4561 2612 1234 5467",
				amount: 5_000_00,
			},
			wantTotal: 5_025_00,
			wantErr:   nil,
		},
		{
			name: "№4 Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Maestro",
							FirstNameOwner: "Bob",
							LastNameOwner:  "Down",
							Balance:        0,
							Currency:       "USD",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5106 2146 1200 5461",
				to:     "4561 2612 1234 5467",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantErr:   ErrNotErrNotEnoughMoney,
		},
		{
			name: "№5 Карта своего банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             333,
							Issuer:         "VISA",
							FirstNameOwner: "Rob",
							LastNameOwner:  "Chock",
							Balance:        10_500_00,
							Currency:       "USD",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5106 2146 1200 5461",
				to:     "4561 2612 1234 5467",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantErr:   nil,
		},
		{
			name: " №6 Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards: []*card.Card{
						{
							Id:             333,
							Issuer:         "VISA",
							FirstNameOwner: "Rob",
							LastNameOwner:  "Chock",
							Balance:        10_500_00,
							Currency:       "USD",
							Number:         "5106 2146 1200 5461",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "4561 2612 1234 5467",
				to:     "5106 2146 1200 5461",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantErr:   nil,
		},
		{
			name: "№7 Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards:        []*card.Card{},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "4561 2612 1234 5467",
				to:     "4561 2646 1200 5608",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantErr:   nil,
		},
		{
			name: "№8 Карта чужого банка -> Карта чужого банка (Неверный номер карты from)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards:        []*card.Card{},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "6666 2612 1234 5467",
				to:     "4561 2646 1200 5608",
				amount: 10_000_00,
			},
			wantTotal: 0,
			wantErr:   ErrInvalidCardNumber,
		},
		{
			name: "№9 Карта чужого банка -> Карта чужого банка (Неверный номер карты to)",
			fields: fields{
				CardSvc: &card.Service{
					BankName:     "Tinkoff",
					NumberPrefix: "5106 21",
					Cards:        []*card.Card{},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "4561 2646 1200 5608",
				to:     "6666 2612 1234 5467",
				amount: 10_000_00,
			},
			wantTotal: 0,
			wantErr:   ErrInvalidCardNumber,
		},
	}
	for _, tt := range tests {
		service := &Service{
			CardSvc:             tt.fields.CardSvc,
			Commission:          tt.fields.Commission,
			MinCommissionAmount: tt.fields.MinCommissionAmount,
		}
		gotTotal, gotError := service.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotError != tt.wantErr {
			t.Errorf("\n test: %v \n Card2Card() gotError = %v, want %v", tt.name, gotError, tt.wantErr)
			return
		}
		if gotTotal != tt.wantTotal {
			t.Errorf("\n test: %v \n Card2Card() gotTotal = %v, want %v", tt.name, gotTotal, tt.wantTotal)
		}
	}
}
