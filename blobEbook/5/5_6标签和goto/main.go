package main

import (
	"fmt"

)


func main() {
	// 
	// 	a := 1
	// 	goto TARGET  // compile error 之前有声明 变量 编译失败
	// 	b := 9
	// TARGET:  
	// 	b += a
	// 	fmt.Printf("a is %v *** b is %v", a, b)
	// i := 0
	// for { //since there are no checks, this is an infinite loop
	// 	if i >= 3 { break }
	// 	//break out of this for loop when this condition is met
	// 	fmt.Println("Value of i is:", i)
	// 	i++
	// }
	// fmt.Println("A statement just after for loop.")
	 test333()
	// q:=Pop(AAA{2,3,1000})
	// fmt.Println(q)
}
var b =&A{name:"111"}
var c A
func  test333()  {

	DoSomethinga(&A{name:"li"})
}
type A struct{
	name string
}
// 命名形参 和 使用的类型应一致
func DoSomethinga(a *A) {
	b = a
	fmt.Println(a,"1212",b)
}

// func DoSomethingb(a A) {
// 	c = &a
// 	c=A{name:"xqxq"}
// }
// func test1(){
// 	for i := 0; i<7 ; i++ {
// 		if i%2 == 0 { continue }
// 		fmt.Println("Odd:", i)
// 	}
// }
// type AAA []int

// func Pop(st AAA) int {
//     v := 0
//     for ix := len(st) - 1; ix >= 0; ix-- {
// 	    fmt.Println(st[ix],v,"1111")
//         if v = st[ix]; v != 0 {
//             // st[ix] = 0  
//             // return v  直接返回
//         }
// 	}
// 	return v
// }  