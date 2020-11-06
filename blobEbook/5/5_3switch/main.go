package main

import "fmt"

func init()  {
	
	num1:=0
	fmt.Printf("%d\n",num1)
}
func switch2(){
	num1:=7;
	switch {
	case num1<0:
		fmt.Println("数字小于0")
	case num1>0&&num1<10:
		fmt.Println("between 0 and 10")
	default:
		fmt.Println("数字10")
		
	}
}
func main(){
	switch2()
	num1:=100
	switch  num1{
		case 98,99:
			fmt.Println("98,99")
		case 100:
			fmt.Println("100")
		default:
			fmt.Println("不等于98or100")
	}
}