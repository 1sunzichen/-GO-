package main

import (
	"fmt"
	"time"
)

func main(){//并发
 	for i:=0;i<10;i++{
 		go func(i int){
 			for{
 				fmt.Printf("Hello goroutine %d\n",
 					i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
 }