package main

import "fmt"

func main(){
	//常量在编译下来就确认的值 iota 自动递增
	const (
		iota1 =iota//0
		iota2 ="111"//1
		iota3 =iota//2
	)
	const iota4=iota//0
	fmt.Print(iota1,iota2,iota3,iota4)
}
