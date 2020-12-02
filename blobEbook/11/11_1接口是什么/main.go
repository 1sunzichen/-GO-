package main

import "fmt"

type stockPosition struct{
	ticker string
	sharePrice float32 
	count float32
}

func (s stockPosition)getValue()float32{
	return s.sharePrice*s.count
}
type car struct{
	make string
	model string
	price float32
}

func (c car)getValue()float32{
	return c.price
}

type valuable interface{
	getValue() float32
}

func showValue(asset valuable){
	fmt.Printf("Value of the asset is %f\n",
asset.getValue())
}
func main(){
	// var o valuable=stockPosition{"GooD",5555.3,10}
	// showValue(o)
	// o=car{"BENZ","G",9099990000.0}
	// showValue(o)
	test()
}


type Simple struct{
	num int
}
func (s Simple)Get()int{
	return s.num
}
func (s *Simple)Set(i int){
	s.num=i
}
type  Simpler interface{
	Get() int
	Set(int)
}
func getorset(Si Simpler){
	fmt.Println(Si.Get());
	Si.Set(9)
	fmt.Println(Si.Get());
}

func test(){
   var s Simpler=&Simple{7}
   getorset(s)
}