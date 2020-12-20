package main

import "fmt"

func main(){
	/*
		int8 1
		int16 2
		int32 4
	   int64 8

	uint8 1
	uint16 2
	uint32 4
	*/
	var int1=100
	var int3 int32
   int3=int32(int1)
   fmt.Printf("%T %T %v",int3,int1,int3)
}
