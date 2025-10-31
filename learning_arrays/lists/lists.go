package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)
	prices = append(prices, 5.99)
	fmt.Println("prices", prices)
	prices = prices[1:] //i.e. like tail
	fmt.Println("like tail", prices)
	fmt.Println("like head", prices[:1]) //i.e. head

	discountPrices := []float64{100.99, 80.99, 12.34}
	prices = append(prices, discountPrices...) // use ... notation to add another array when using append
	fmt.Println(prices)
}

//
//func main() {
//
//	var productNames [4]string = [4]string{"blah"}
//
//	prices := [4]float64{1.1, 1.2, 2.4, 5.6}
//	fmt.Println(prices)
//	productNames[1] = "huhwuh"
//	productNames[2] = "huh"
//	fmt.Println(productNames)
//	fmt.Println(prices[0])
//
//	//slices
//	featuredPrices := prices[1:3] // i.e. [1.2 2.4]
//
//	fmt.Println(featuredPrices)
//
//	fmt.Println(prices[:2]) //i.e. start at first index, ot the end would be something like [1:]
//
//	fmt.Println(len(featuredPrices), cap(featuredPrices))
//	fmt.Println(len(productNames), cap(productNames))
//
//}
