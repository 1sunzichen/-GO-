package main

import (
	"fmt"
	"strconv"
	"os"
)
func main(){
	orig:="ABC"
	var newS string
	fmt.Printf("The size of ints is:%d\n",strconv.IntSize)
	an,err:=strconv.Atoi(orig)// 字符串转数字
	if err !=nil{
		fmt.Printf("你%s\n",orig)
		os.Exit(1) //如果我们想要在错误发生的同时终止程序的运行，我们可以使用 os 包的 Exit 函数：
	}
	fmt.Printf("The integer is %d\n",an)
	an=an+5
	newS=strconv.Itoa(an)//数字 转 字符串
	fmt.Printf("The new string is: %s\n", newS)
}