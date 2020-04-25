package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 5 + 12i
	fmt.Println(cmplx.Abs(c))
	//欧拉公式
	fmt.Println(cmplx.Exp(1i*math.Pi)+1, cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func ooo() {
	fmt.Printf("%.0f\n",
		cmplx.Exp(1i*math.Pi)+1)

}

//枚举
func enums() {
	const (
		java = iota
		_
		python
		golang
	)
	const (
		//1<<  1*2的iota*10次方
		//iota 从0算起 0，1，2
		b = 1 << (10 * iota)
		kb
	)
	fmt.Println(java, python, golang)
	fmt.Println(b, kb)
}

func main() {
	//ooo()
	//euler()
	triangle()
	consts()
	enums()
}
