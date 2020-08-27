package main

import (
	"fmt"
	"sync"
)
func worker(id int,c chan int,wg *sync.WaitGroup ) {
	//for {
	//		n,ok:=<-c
	//		if !ok{
	//			break
	//		}
	//	fmt.Printf("Worker %v received %d\n", id, n)
	//}
	for n:=range c{

		fmt.Printf("Worker %v received %d\n", id, n)
		//go func(){
		//	//通信 共享内存
		//	done <-true
		//
		//}()
		wg.Done()
	}
}

type workerx struct {
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}
//func worker(id int,())
func createworker(id int,wg *sync.WaitGroup) workerx{
	w:=workerx{
		in:make(chan int),
		//done:make(chan bool),
		wg:wg,
	}
	// c:=make(chan int)
	//f := func() {
	//	for {
	//
	//		fmt.Printf("Worker %v received %d\n", id, <-c)
	//	}
	//}
	go worker(id,w.in,w.wg)

	return w
}
func chanDemo(){
	//var channels [10]chan<- int
	var workers [10]workerx
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		workers[i]=createworker(i,&wg)//

	}
	wg.Add(10)
	//channel 一等公民
	//for i,wo:=range workers{
	//	wo.in<-'a'+i
	//	//<-workers[i].done
	//}
	for i,wo:=range workers{
		wo.in<-'A'+i
		//<-workers[i].done
	}
	wg.Wait()
	//for i:=0;i<10 ;i++{
	//	workers[i].in<-'A'+i
	//	//<-workers[i].done
	//}
	//for i:=0;i<10 ;i++{
	//	workers[i].in<-'B'+i
	//	<-workers[i].done
	//}
	//for i:=0;i<10 ;i++{
	//	workers[i].in<-'b'+i
	//	<-workers[i].done
	//}

	//for _,workerx:=range  workers{
	//	<-workerx.done
	//}
	//c<-1
	//c<-2
	//time.Sleep(time.Millisecond)

}
func  main()  {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
