package main

import (
	"fmt"
)

//使用close 关闭程序
func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ok = true
	var i int
	ch := make(chan int)
	go tel(ch)
	for ok {
		if i, ok = <-ch; ok {
			fmt.Println("ok", ok, i)
		}
	}
}
