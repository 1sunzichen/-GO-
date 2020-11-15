package main

import (
"fmt"
"math"
"math/big"
)

func main() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	//               分子          分母
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	// 乘           加          乘
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
	a:=big.NewRat(1,2);
	fmt.Printf("%v",a)
}
