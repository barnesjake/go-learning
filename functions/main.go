package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15, 40, -5)
	numbers := []int{1, 10, 15}
	anotherSum := sumup(0, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// / ...  is syntax for var arg list (go calls it variadic function), collect all parameter
func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
