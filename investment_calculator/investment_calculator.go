package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const inflationRate float64 = 2.5
const resultsFileName = "results.txt"

func getResultsFromFile() (float64, error) {
	data, err := os.ReadFile(resultsFileName)
	if err != nil {
		return 1000, errors.New("could not find results file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("could not parse value in results file")
	}

	return balance, nil
}

func writeResultsToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(resultsFileName, []byte(balanceText), 0644)
}

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	var futureValue float64 = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Println("Future value:", futureValue)
	//fmt.Printf("Future real value: %.2f", futureRealValue)

	formattedFv := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRfv := fmt.Sprintf("Future real value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFv, formattedRfv)
	fmt.Println("Writing results to file...")
	writeResultsToFile(futureRealValue)

}

//	func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
//		fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
//		rfv := fv / math.Pow(1+inflationRate/100, years)
//		return fv, rfv
//	}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
