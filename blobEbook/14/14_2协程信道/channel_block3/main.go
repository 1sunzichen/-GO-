package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go setData(ch1)
	fmt.Println("sendding", 10)
	ch1 <- "10"
	fmt.Println("send", 10)
}

func setData(ch1 chan string) {
	time.Sleep(1e9 * 15)

	x := <-ch1
	fmt.Println(x, "received")
}
