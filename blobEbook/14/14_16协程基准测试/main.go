package main

import(
	"fmt"
	"testing"
)

func main(){
	fmt.Println("Sync",testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("Buffer",testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B){
	ch:=make(chan int)
	go func(){
		for i:=0;i<b.N;i++{
			ch<- i
		}
		close(ch)
	}()
	for range ch{

	}
}
func  BenchmarkChannelBuffered(b *testing.B)  {
	ch:=make(chan int,128)
	go func(){
		for i:=0;i<b.N;i++{
			ch<-i
		}
		close(ch)
	}()
	for range ch{

	}
}
//意味着循环执行了 b.N 次，每次循环花费 xx 纳秒(ns)。