package main

import "fmt"

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	//website1 := map[string]int{
	//	"Google": 1,
	//	"AWS":    2,
	//}
	fmt.Println(websites)        //map[AWS:https://aws.com Google:https://google.com]
	fmt.Println(websites["AWS"]) //https://aws.com

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites) //map[AWS:https://aws.com Google:https://google.com LinkedIn:https://www.linkedin.com]

	delete(websites, "Google")
	fmt.Println(websites) //map[AWS:https://aws.com LinkedIn:https://www.linkedin.com]

	//make - arrays
	//userNames := make([]string, 2) //make an array behind the scenes with 2 slots, useful if you want to do userNames[0] = "blah"
	userNames := make([]string, 2, 5) //max 5 items, 2 empty slots. Go doesn't have to go and allocate new memory, making append more efficient. You can still add more than the max though, but then it starts allocating new memory

	userNames = append(userNames, "John")
	userNames = append(userNames, "Bob")
	fmt.Println(userNames) // [  John Bob]
	
	for key := range websites {
		fmt.Println(key)
	}

	//make - maps
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 1.0
	courseRatings["scala"] = 5.0
	fmt.Println(courseRatings) // map[go:4.7 react:1 scala:5]

	courseRatings.output()

	//for loop over map
	for key, value := range courseRatings {
		fmt.Println("Key:", key, "Value:", value)
	}

}
