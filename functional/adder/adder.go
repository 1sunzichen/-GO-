package main

import "fmt"

//函數式編程 閉包
func adder() func(int) int{
	sum:=0
	return func(v int) int{
		sum+=v
		return sum
	}
}

type iAdder func(int)(int,iAdder)
func adder2(base int) iAdder{
	return func(v int)(int,iAdder){
		return base+v,adder2(base+v)
	}
}
func main(){
	a:=adder()
	a2:=adder2(0)
	for i:=0;i<10;i++{
		var s2 int
		s2,a2=a2(i)
		fmt.Printf("0+1+...+%v=%v,%d\n",i,a(i),s2)
	}
}
