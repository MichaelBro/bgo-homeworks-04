package main

import (
	"bgo-homeworks-04/pkg/card"
	"bgo-homeworks-04/pkg/transfer"
	"fmt"
)

func main() {
	service := card.NewService("Main Bank")

	visa := service.IssueCard(
		1,
		"Visa",
		"Michael",
		"Bro",
		50_000_00,
		"RUB",
		"4800 8000 9000 1234",
	)

	master := service.IssueCard(
		2,
		"MasterCard",
		"Michael",
		"Bro",
		12_000_00,
		"RUB",
		"5500 8000 1234 9876",
	)

	fmt.Println(visa)
	fmt.Println(master)

	serviceTransfer := transfer.NewService(service, 0.5, 10)

	serviceTransfer.Card2Card(visa.Number, master.Number, 10_000_00)

	fmt.Println(visa)
	fmt.Println(master)

}
