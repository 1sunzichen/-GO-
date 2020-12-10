package main


import "fmt"

func c(x int) func(k int)int{
	var a int
	a=x
	return func(k int)int{
		b:=a
		// a=k
		return b * k
	}
}
func main(){
	f:=c(1)
	fmt.Println(f(2))
	fmt.Println(f(3))
}