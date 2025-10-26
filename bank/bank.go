package main

import (
	"bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		panic("cannot continue, sorry!")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int
		fmt.Print("\nYour choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice:", choice)

		switch choice {

		case 1:
			fmt.Printf("Your account balance: %.2f\n", accountBalance)

		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Your deposit amount must be positive")
				continue
			}

			//accountBalance = accountBalance + depositAmount
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
		case 3:
			fmt.Print("How much do you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Your withdrawal amount must be positive")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Your withdrawal amount cannot be more than your account balance!")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)

		case 4:
			fmt.Println("Goodbye!")
			break
		}
	}
}
