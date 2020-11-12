package main

import (
	"fmt"
	"sort"
)

func main(){

	// test2()
	sortmap()
}

func sortmap(){
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
								"delta": 87, "echo": 56, "foxtrot": 12,
								"golf": 34, "hotel": 16, "indio": 87,
								"juliet": 65, "kili": 43, "lima": 98}
	)
	fmt.Println("unsorted:")
	// for k, v := range barVal {
	// 	fmt.Printf("Key: %v, Value: %v / ", k, v)
	// }
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	fmt.Println(keys,"keys")
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

func test2(){
	// Version A:
	items := make([]map[int]int, 5)
	for i:= range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

func weektest(){
	map2:=map[string]string{
		"7":"Sunday",
		"1":"Monday",
		"2":"Tuesday",
		"3":"Wednesday",
		"4":"Thursday",
		"5":"Friday",
		"6":"Saturday",
	}
	for key,value:=range map2{
		if(key=="7"||key=="6"){
			fmt.Printf("%v %v\n",key,value)
		}
		if(value=="Thursday"){
			fmt.Printf("%v %v\n",key,value)
		}
	}
}
func init(){
	map1:=make(map[int]float32)
	map1[0]=1.0
	map1[04]=1.0
	map1[03]=1.0
	map1[02]=1.0
	for key,value:=range map1{
		fmt.Printf("%d %f\n",key,value)
	}

}