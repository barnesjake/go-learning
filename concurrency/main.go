package main

import (
	"fmt"
	"time"
)

func main() {
	//go greet("Nice to meet you")
	//go greet("Nice to meet you 1")
	//done := make(chan bool)
	//go greetSlow("Nice to meet you 2", done)
	//go greet("Nice to meet you 3")
	//<-done // Arrow describes the direction of dataflow. Blocks, i.e. go waits for the data to come out of the channel

	//done := make(chan bool)
	//go greet("Nice to meet you", done)
	//go greet("Nice to meet you 1", done)
	//go greetSlow("Nice to meet you 2", done)
	//go greet("Nice to meet you 3", done)
	//<-done //waits for the first response in channel, so can cause race condition - issue with reusing channels. you could wait 4 times in this example but that's not great

	//dones := make([]chan bool, 4)
	//dones[0] = make(chan bool)
	//go greet("Nice to meet you", dones[0])
	//dones[1] = make(chan bool)
	//go greet("Nice to meet you 1", dones[1])
	//dones[2] = make(chan bool)
	//go greetSlow("Nice to meet you 2", dones[2])
	//dones[3] = make(chan bool)
	//go greet("Nice to meet you 3", dones[3])
	//for _, done := range dones {
	//	<-done
	//}

	done := make(chan bool)

	go greet("Nice to meet you", done)
	go greet("Nice to meet you 1", done)
	go greetSlow("Nice to meet you 2", done)
	go greet("Nice to meet you 3", done)
	//use the channel itself as the lim on the for loop
	//for doneChan := range done {
	//	fmt.Println(doneChan)
	//}
	//or just block and wait doing nothing with them
	for range done {
	}

}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func greetSlow(phrase string, doneChan chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	//only use when you know it's not needed anymore
	close(doneChan)
}
