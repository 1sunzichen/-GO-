package main

import (
	"fmt"
	"os"
)

//接受 通道后 ，关闭程序
func tel(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i

	}
	quit <- true
}

func main() {
	ok := true
	ch := make(chan int)
	quit := make(chan bool)
	go tel(ch, quit)
	for ok {
		select {
		case i := <-ch:
			fmt.Printf("counter%d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}
