package main

import "fmt"

func main() {

	var result = add(1, 2)
	fmt.Println(result)
}

// func add[T any](a, b T) T {
func add[T int | float64 | string](a, b T) T {
	return a + b
}
