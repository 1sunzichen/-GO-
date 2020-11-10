package main

import (
	"fmt"
	
)
func main(){
	// test1()
	// maketest()
	makenew()
	test73()
}

// func Append(slice []byte,b []byte)[]byte{
// 	len1:=len(slice)+len(b)
// 	slice2:=make([]byte,len1)
// 	slice2[0:len(slice)]=slice
// 	slice2[len(slice):]=b
// 	return slice2
// }

func test73()  {
	// s:=make([]byte,5)//类型 长度 
	// a:=s[2:4]
	// fmt.Println(len(s),cap(s),len(a),cap(a))
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	s2[1] = 't'

	fmt.Printf("%v %v",s2,s1)
}
func makenew(){
	p := new([]int)
	v := make([]int, 0)
	fmt.Println(p,v)

}
func maketest(){
	var slice1 []int=make([]int,10)
	for i := 0; i < len(slice1); i++ {
		slice1[i]=5*i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n",i,slice1[i])
	}
	fmt.Println(len(slice1),cap(slice1))
}

func test1(){
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4],b[:2],b[2:],b[:])
}
func example(){
	var arr1 [7]int
	var slice1 []int=arr1[2:6]//索引2-索引6-1
	for i := 0; i < len(arr1); i++ {
		arr1[i]=i
	}
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("%d",)
	// }
	fmt.Print(arr1)
	for i,v:=range slice1{
		fmt.Printf("%d ==%d\n",i,v)
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	//切片的长度 + 数组除切片之后的长度
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice
	slice1 = slice1[0:5]//0-3 +1 cap  切片不够 用原数组的
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

