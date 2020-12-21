package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	// b := <-ch1
	go pump2(ch1)
	fmt.Print(<-ch1)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func pump2(ch chan int) {

	for {
		input := <-ch
		fmt.Print(input)
	}
}
