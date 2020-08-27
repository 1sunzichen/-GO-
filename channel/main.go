package main

import (
	"fmt"
	"time"
)
func worker(id int,c chan int) {
	//for {
	//		n,ok:=<-c
	//		if !ok{
	//			break
	//		}
	//	fmt.Printf("Worker %v received %d\n", id, n)
	//}
	for n:=range c{//读

		fmt.Printf("Worker %v received %d\n", id, n)

	}
}
//func worker(id int,())
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

func bufferedChannel(){
	c:=make(chan int,3)// 3缓冲区
	go worker(0,c);
	c<-'a'
	c<-'b'
	c<-'c'
	c<-'d'
	time.Sleep(time.Millisecond)
}
func channelClose(){
	c:=make(chan int)// 3缓冲区
	go worker(0,c);
	c<-'a' //写 先执行 写
	c<-'b'
	c<-'c'
	c<-'d'
	//close(c)
	//time.Sleep(time.Millisecond)
}
func chanDemo(){
	var channels [10]chan<- int
	for i:=0;i<10;i++{
		channels[i]=createworker(i)

	}
	//channel 一等公民
	for i:=0;i<10 ;i++{
		channels[i]<-'a'+i
	}
	for i:=0;i<10 ;i++{
		channels[i]<-'A'+i
	}
	for i:=0;i<10 ;i++{
		channels[i]<-'B'+i
	}
	for i:=0;i<10 ;i++{
		channels[i]<-'b'+i
	}
	//c<-1
	//c<-2
	time.Sleep(time.Millisecond)

}
func  main()  {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
