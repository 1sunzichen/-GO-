package even

import "fmt"

//偶数
func Even(i int) bool {
	fmt.Print(i%2 == 0, i)

	return i%2 == 0
}

func Odd(i int) bool {
	return i%2 != 0
}
