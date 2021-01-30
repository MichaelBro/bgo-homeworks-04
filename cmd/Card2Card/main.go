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
		"5106 2100 9000 1234",
	)

	master := service.IssueCard(
		2,
		"MasterCard",
		"Michael",
		"Bro",
		0,
		"RUB",
		"5500 8000 1234 9876",
	)

	fmt.Println(visa)
	fmt.Println(master)

	serviceTransfer := transfer.NewService(service, 0.5, 10)

	total, err := serviceTransfer.Card2Card(visa.Number, master.Number, 5_000_00)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(total)
	fmt.Println(visa)
	fmt.Println(master)

}
