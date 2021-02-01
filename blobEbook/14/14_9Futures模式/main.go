package main

import "fmt"

type Any interface{}

func main() {
	A, B := InverseProduct(5, 6)
	fmt.Println(A, B, "number")
}
func InverseProduct(a Any, b Any) (Any, Any) {
	a_inv_future := InverseFuture(a) // start as a goroutine
	b_inv_future := InverseFuture(b) // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return a_inv, b_inv
}
func InverseFuture(a Any) chan Any {
	future := make(chan Any)
	go func() {
		future <- a.(int) * 2
	}()
	return future
}
