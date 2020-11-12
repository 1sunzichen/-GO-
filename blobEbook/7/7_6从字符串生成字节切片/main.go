package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c", i, c)
	}
	//str:="一颗呀小白杨长在哨所旁"
	//substr :=  str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(mapFunc(multiplication, []int{1, 2, 3, 4, 1, 2, 1, 21}))

}
func multiplication(i int) int {
	return i * 10
}

func mapFunc(fn func(int) int, arr []int) []int {
	var arrs []int
	for _, v := range arr {
		arrs = append(arrs, multiplication(v))
	}
	return arrs
}

func mapp(arr []int) []int {
	//冒泡排序
	for i, _ := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}

func test715(s string) (string, string) {
	var str []byte
	for _, v := range s {
		if strings.Contains(string(str), string(v)) == false {
			str = append(str, byte(v))
		}
	}
	return string(str), s
}

func test1(str string, i int) (string, string) {
	arr := str[0:i]
	arr2 := str[i:]
	return arr, arr2

}
func nf(c string) string {
	s := []byte(c)
	for i, v := range s {
		if i < int(len(s)/2) {

			t := s[len(s)-1-i]
			s[(len(s) - 1 - i)] = v
			s[i] = t

		}
		break
	}
	return string(s)
}
