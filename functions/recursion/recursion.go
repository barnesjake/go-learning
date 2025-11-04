package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(n int) int {
	//for recursive func, remember to define exit condition to avoid infinite loop
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	//result := 1
	//
	//for i := 1; i <= n; i++ {
	//	result *= i
	//}
	//
	//return result
}
