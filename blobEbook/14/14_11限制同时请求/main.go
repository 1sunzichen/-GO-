package main

import (
	"fmt"
)
const MAXREQS=20

var sem=make(chan int,MAXREQS)

type Request struct{
	a,b int
	replyc chan int
}

func process(r *Request){
	//do something
	fmt.Println(r)
}

func handle(r *Request){
	sem<-1
	process(r)
	<-sem
}


func server(service chan *Request){
	for{
		request:=<-service
		go handle(request)
	}

}

func main(){
	service:=make(chan *Request)
	go server(service)
}