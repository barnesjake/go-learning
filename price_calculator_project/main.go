package main

import (
	"fmt"
	"price_calculator_project/filemanager"
	"price_calculator_project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		//if err != nil {
		//	fmt.Println("Could not process job")
		//	fmt.Println(err)
		//}
	}
	//for _, taxRate := range taxRates {}
	for index := range taxRates {

		// select lets you wait on multiple channel operations
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

}
