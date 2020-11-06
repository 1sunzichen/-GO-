package main
import (
	"fmt"
	"math"
	"errors"
)

func calculating(a int,b int)[]int{
	return []int{a+b,a-b,a*b}
}
func calculating2(a int,b int){
	fmt.Println(a+b,a-b,a*b)
}

func MySqrt(a float64)float64{
	
	return math.Sqrt(a)
}
func MySqrt2(a float64)(b float64,err error){
	if(a<0){
		b = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	}else{

		b=math.Sqrt(a)
	}
	return 
}
func Multiply(a, b int, reply *int) {
    *reply = a * b
}
func outsidevariable(){
	n:=0
	reply := &n
    Multiply(10, 5, reply)
    fmt.Println("Multiply:", *reply) // Multiply: 50

}
func main(){
	a:=calculating(3,5)
	fmt.Println(a)
	c,_:=MySqrt2(-8)
	fmt.Println(c)
	outsidevariable()
}