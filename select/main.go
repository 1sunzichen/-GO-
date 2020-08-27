package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out:=make(chan int)
	go func(){
		i:=0
		for{
			time.Sleep(
				time.Duration(rand.Intn(1500))*time.Millisecond)
			out<-i //送分童子
			i++
		}
	}()
	return out
}
func worker(id int,c chan int) {
	for n:=range c{//读
		fmt.Printf("Worker %v received %d\n", id, n)

	}
}
func createworker(id int) chan<- int{
	c:=make(chan int)
	//f := func() {
	//	for {
	//
	//		fmt.Printf("Worker %v received %d\n", id, <-c)
	//	}
	//}
	go worker(id,c)

	return c
}
func main(){
	var c1,c2=generator(),generator()
    var w=createworker(0)
	n:=0
	hasValue:=false
	for{
		var activeWork chan<- int
		if hasValue{
			activeWork=w
		}
		select {
			case n=<-c1:
				hasValue=true
			case n=<-c2:
				hasValue=true

			case activeWork <-n:
			//default:
			//	fmt.Println("mmmmm")
		}
	}

}
