package main
import (
	"runtime"
	"fmt"
)
const NCPU = 4
func DoAll(){
	sem:=make(chan int,NCPU)
	for i:=0;i<NCPU;i++{
		go DoPart(sem)
	}
	for i:=0;i<NCPU;i++{
		<-sem
	}
	fmt.Print("AllDone")
}
func DoPart(sem chan int){
	sem<-1
}
func main(){
	runtime.GOMAXPROCS(NCPU)
	DoAll()
}