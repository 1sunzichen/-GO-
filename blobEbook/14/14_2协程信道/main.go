package main

import (
	"fmt"
	"time"
)

//如果我们不在 main() 中等待，协程会随着程序的结束而消亡。
func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)

}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}
