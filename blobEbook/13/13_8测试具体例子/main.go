package main

import (
	"123go/blobEbook/13/13_8测试具体例子/even"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
