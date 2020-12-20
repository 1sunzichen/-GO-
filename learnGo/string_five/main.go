package main

import "fmt"

// byte uint8
//rune int8
//字符=》ascii 码值 =》二进制
func main(){
	 var str1="NIHAO WOSHI AIKE"
	 for _,st:=range str1{
	 	fmt.Printf(
	 		"  当前%s ,  编码%d 值=%c    类型=%T,\n",
			str1,st,st,st)
	 }
}
