package main

import (
	"fmt"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	// 13 转成2进制
	//13  6  3  1   到0不在除
	// 1	0  1  1
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}
func main() {
	fmt.Println(
		converToBin(3),
		converToBin(13),
	)
}
