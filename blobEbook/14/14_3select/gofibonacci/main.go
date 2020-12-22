package main

import (
	"fmt"
	"time"
)

func main() {
	term := 25
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibnterms(term, c)
	for {
		if result, ok := <-c; ok {
			fmt.Printf("%d %d", i, result)
			i++
		} else {
			end := time.Now()
			delta := end.Sub(start)
			fmt.Printf("%s\n", delta)
			// os.Exit(0)
		}

	}

}
func fibnterms(term int, c chan int) {
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
