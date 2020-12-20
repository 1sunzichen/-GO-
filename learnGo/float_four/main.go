package main

import "fmt"

func main(){
	// float32  4个字节
	// float64  8个字节
	var float3 float32=3.14
	var float4 float64=3.15
	float4=float64(float3)
	fmt.Printf("%T %v",float4,float3)
	// 复数  实数 虚数
	//complex64 32实数+32虚数 128 64实数 + 64虚数
	var complex1 complex64
	complex1=3.14+12i
	complex2:=complex(3.14,12)
   fmt.Printf("%T %v",complex1,complex2)

}
