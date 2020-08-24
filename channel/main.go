package main

import (
	"fmt"
	"time"
)
func createworker(id int) chan<- int{
	c:=make(chan int)
	go func(){
		for{

			fmt.Printf("Worker %v received %d\n",id,<-c)
		}
	}()

	return c
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
	chanDemo()
}
