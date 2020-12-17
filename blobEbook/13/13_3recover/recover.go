package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}
func main() {
	fmt.Println("Start")
	test()
	fmt.Println("End")
}
func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r \n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}
