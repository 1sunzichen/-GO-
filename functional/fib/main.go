package main

import (
	"123go/functional/fib/fibonacci"
	"bufio"
	"fmt"
	"io"
	"strings"
)

//func fibonacci() intGen{
//	a,s:=0,1
//	return func() int {
//		// 下一次啊 a 的位置 是 S   s的位置是和
//		a,s=s,a+s
//		//fmt.Println(a)
//		return a
//	}
//}
type intGen func() int

func (g intGen) Read(p []byte)(n int,err error){
	next:=g()
	if next>10000{
		return 0,io.EOF
	}
	// TODO:
	s:=fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}
func printFileContents(reader io.Reader){
	scanner:=bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func main(){
	f:=fibonacci.Fibonacci()
	//f()
	//f()
	//f()
	//f()
	//f()
	printFileContents(f)
}
