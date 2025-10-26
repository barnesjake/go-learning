package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
//  1. Validate user input
//     => show error message & exit if invalid input is provided
//     - no negative numbers
//     - not 0/zero
//  2. store calculated results into file

func main() {

	revenue, err := getUserInput("Revenue: ")
	errorAndPanic(err)
	expenses, err := getUserInput("Expenses: ")
	errorAndPanic(err)
	taxRate, err := getUserInput("Tax Rate (%): ")
	errorAndPanic(err)

	ebt, profit, earningsRatio := calculateEbtProfitAndEarningsRatio(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Earnings ratio: %.1f\n", earningsRatio)

	storeResults(ebt, profit, earningsRatio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0.0 {
		return 0, errors.New(infoText + "cannot be zero or negative")
	}

	return userInput, nil
}

func errorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateEbtProfitAndEarningsRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {

	//earnings before tax
	var ebt float64 = revenue - expenses
	//earnings after tax
	var profit float64 = ebt * (1 - taxRate/100)
	//ebt/profit
	var earningsRatio float64 = ebt / profit

	return ebt, profit, earningsRatio
}
