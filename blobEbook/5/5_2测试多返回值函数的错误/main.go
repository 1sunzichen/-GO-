package main

import (
	"fmt"
	"strconv"
)
func main(){
	orig:="ABC"
	var newS string
	fmt.Printf("The size of ints is:%d\n",strconv.IntSize)
	an,err:=strconv.Atoi(orig)// 字符串转数字
	if err !=nil{
		fmt.Printf("你%s\n",orig)
		return 
	}
	fmt.Printf("The integer is %d\n",an)
	an=an+5
	newS=strconv.Itoa(an)//数字 转 字符串
	fmt.Printf("The new string is: %s\n", newS)
}