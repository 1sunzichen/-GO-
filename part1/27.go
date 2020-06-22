package main

import (
	// "fmt"
	"strconv"
)

func convertToBin(n int) string{
	if n==0 {
		return "0";
	}
	result:=""
	for ;n>0;n/=2{
		lsb:=n%2
		result=strconv.Itoa(lsb)+result
	}
	return result;
}


// func main(){
// 	fmt.Println(
// 		convertToBin(5),//101
// 		convertToBin(0),//1101 
// 	)
// }