package main

import "fmt"

func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println(`
	1. Check balance
	2. Deposit money
	3. Withdraw money
	4. Exit`)
}
