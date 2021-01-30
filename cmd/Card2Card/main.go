package main

import (
	"bgo-homeworks-04/pkg/card"
	"bgo-homeworks-04/pkg/transfer"
	"fmt"
)

func main() {
	service := card.NewService("Main Bank", "5106 21")

	visa := service.IssueCard(
		1,
		"Visa",
		"Michael",
		"Bro",
		16_125_99,
		"RUB",
		"5106 2112 1234 5461",
	)

	master := service.IssueCard(
		2,
		"MasterCard",
		"Michael",
		"Bro",
		0,
		"RUB",
		"4561 2612 1234 5467",
	)

	fmt.Println(visa)
	fmt.Println(transfer.IsValid(visa.Number))

	fmt.Println(master)
	fmt.Println(transfer.IsValid(master.Number))

	inValidNumber := "4561 2612 1234 5464"
	validNumber := "4561 2612 1234 5467"

	fmt.Println(transfer.IsValid(inValidNumber))
	fmt.Println(transfer.IsValid(validNumber))

	serviceTransfer := transfer.NewService(service, 0.5, 10)

	total, err := serviceTransfer.Card2Card(visa.Number, master.Number, 5_000_00)

	if err != nil {
		fmt.Println(err)
		fmt.Println(total)
	}

	fmt.Println(total)
	fmt.Println(visa)
	fmt.Println(master)

}
