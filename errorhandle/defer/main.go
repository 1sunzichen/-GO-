package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer(){
	defer fmt.Println(1)//棧 先進後出
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error")

}

func writeFile(filename string){
	file,err:=os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	write:=bufio.NewWriter(file)
	f:=fib.Fib()

}
func main(){
	tryDefer()
}
