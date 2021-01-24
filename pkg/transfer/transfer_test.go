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
		wantOk    bool
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             1,
							Issuer:         "Visa",
							FirstNameOwner: "Michael",
							LastNameOwner:  "Bro",
							Balance:        100_000_00,
							Currency:       "RUB",
							Number:         "4444 9999 9000 1234",
							Icon:           "icon.png",
						},
						{
							Id:             1,
							Issuer:         "Visa",
							FirstNameOwner: "Alex",
							LastNameOwner:  "Storm",
							Balance:        50_000_00,
							Currency:       "RUB",
							Number:         "5555 7777 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "4444 9999 9000 1234",
				to:     "5555 7777 9000 1234",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Visa",
							FirstNameOwner: "Michael",
							LastNameOwner:  "Bro",
							Balance:        5000_00,
							Currency:       "RUB",
							Number:         "5555 9999 9000 1234",
							Icon:           "icon.png",
						},
						{
							Id:             0,
							Issuer:         "Visa",
							FirstNameOwner: "Alex",
							LastNameOwner:  "Storm",
							Balance:        50_000_00,
							Currency:       "RUB",
							Number:         "5555 7777 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "5555 9999 9000 1234",
				to:     "5555 7777 9000 1234",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantOk:    false,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Maestro",
							FirstNameOwner: "Bob",
							LastNameOwner:  "Down",
							Balance:        16_125_99,
							Currency:       "USD",
							Number:         "1234 9999 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "1234 9999 9000 1234",
				to:     "5555 7777 9999 1234",
				amount: 5_000_00,
			},
			wantTotal: 5_025_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             16,
							Issuer:         "Maestro",
							FirstNameOwner: "Bob",
							LastNameOwner:  "Down",
							Balance:        0,
							Currency:       "USD",
							Number:         "1234 9999 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "1234 9999 9000 1234",
				to:     "5555 7777 9999 1234",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantOk:    false,
		},
		{
			name: "Карта своего банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             333,
							Issuer:         "VISA",
							FirstNameOwner: "Rob",
							LastNameOwner:  "Chock",
							Balance:        10_500_00,
							Currency:       "USD",
							Number:         "6666 9999 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "6666 9999 9000 1234",
				to:     "1234 9999 9000 1234",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantOk:    true,
		},
		{
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "Tinkoff",
					Cards: []*card.Card{
						{
							Id:             333,
							Issuer:         "VISA",
							FirstNameOwner: "Rob",
							LastNameOwner:  "Chock",
							Balance:        10_500_00,
							Currency:       "USD",
							Number:         "6666 9999 9000 1234",
							Icon:           "icon.png",
						},
					},
				},
				Commission:          0.5,
				MinCommissionAmount: 10,
			},
			args: args{
				from:   "1234 9999 9000 1234",
				to:     "6666 9999 9000 1234",
				amount: 10_000_00,
			},
			wantTotal: 10_050_00,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		service := &Service{
			CardSvc:             tt.fields.CardSvc,
			Commission:          tt.fields.Commission,
			MinCommissionAmount: tt.fields.MinCommissionAmount,
		}
		gotTotal, gotOk := service.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotTotal != tt.wantTotal {
			t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
		}
		if gotOk != tt.wantOk {
			t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
		}
	}
}
