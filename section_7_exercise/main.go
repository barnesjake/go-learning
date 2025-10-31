package main

import (
	"fmt"
	"section_7_exercise/product"
)

func main() {

	hobbies := [3]string{
		"Cooking",
		"Gardening",
		"Gaming",
	}

	fmt.Println("1) My hobbies:", hobbies)

	fmt.Println("2.1) First element of hobbies:", hobbies[0])
	//var tailOfHobbies = hobbies[1:]
	var tailOfHobbies = [2]string{hobbies[1], hobbies[2]}
	fmt.Println("2.2) Second and third elements of hobbies:", tailOfHobbies)

	sliceV1 := hobbies[0:2]
	sliceV2 := hobbies[:2]
	fmt.Println("3.1) Slice of hobbies:", sliceV1)
	fmt.Println("3.2) Slice of hobbies another way:", sliceV2)

	sliceV3 := append(sliceV1[1:], hobbies[2])
	fmt.Println("4) Re-slice:", sliceV3)

	courseGoals := []string{"Learn GO syntax", "Write a REST API"}
	fmt.Println("5) Course Goals:", courseGoals)

	courseGoals[1] = "Some other goal"
	courseGoals = append(courseGoals, "A third goal of some kind")
	fmt.Println("6) Course Goals modified:", courseGoals)

	listOfProducts := []product.Product{
		product.New("apples", 1, 0.99),
		product.New("oranges", 2, 1.99),
	}
	fmt.Println("7.1) Products:", listOfProducts)

	listOfProducts = append(listOfProducts, product.New("pears", 3, 0.5))
	fmt.Println("7.2) Products modified:", listOfProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
