package main

import (
	"fmt"
	"math/big"
)

func  main()  {
	// result:=0
	var i int64
	for i= 1; i <= 30; i++ {
		// result=tt(i)
		fmt.Println(nn(big.NewInt(i)))
		// fmt.Println(i,result)	
	}	
	// tt(10)
	// fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
}
func fb(n int)(res int){
	if n<=1{

		res=1
	}else{
		res=fb(n-1)+fb(n-2)
	}
	return
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}


func nn(b *big.Int)*big.Int{
	
	if b.Int64()==1{
		return big.NewInt(1)
	}else{
		return b.Mul(b,nn(big.NewInt(b.Int64()-1)))
	}
}
func tt(b int){
	if b<=1{
		fmt.Println(1);
		// re=1
	}else{
		fmt.Println(b);
		tt(b-1)
	}
	// return 
}