package main

import (
	"time"
)

//func printHello(i int)  {
//	for{
//		fmt.Printf("Hello from"+"goroutine %d",i)
//
//	}
//
//}
func main(){
	// var b [11]int
	a:=0
	for i:=0;i<10;i++{
		go func(i int){ //一定时间 可以执行几次 并发执行
			for{
				//b[i]++ //超过 定义的a 的边界
				//runtime.Gosched()
				a++
				//fmt.Printf("Hello from"+"goroutine %d, %v\n",i,b) io 操作 抢占控制权

				//fmt.Printf("Hello from"+"goroutine %d, %v\n",i,a)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
