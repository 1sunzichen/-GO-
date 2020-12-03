package main

import (
	"fmt"
	"123go/blobEbook/11/11_12接口与动态类型/栈方法实现"
)
// 实现泛型
type obj interface{}
type Car struct{
	name string
	age  int
}
func main() {
	// define a generic lambda function mf:
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, isl)
	fmt.Println(res1)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []obj{"0", "1", "2", "3", "4", "5"}
	res2 := mapFunc(mf, ssl)
	fmt.Println(res2)
	for _, v := range res2 {
		fmt.Println(v)
	}
	abc:= stack.Stack{Car{"小明",12},2,3}
	abc.Push(2)
	fmt.Print(abc,"abc",abc.Cap(),abc.Len())

}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		fmt.Print(v,"vvvv")
		result[ix] = mf(v)
	}

	// Equivalent:
	/*
		for ix := 0; ix<len(list); ix++ {
			result[ix] = mf(list[ix])
		}
	*/
	return result
}