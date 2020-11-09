package main

import(
	"fmt"
)

func main(){
  var arr1 [7]int
	for i := 0; i < len(arr1); i++ {
	arr1[i]=i*2	
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index%d %d\n",i,arr1[i])
	}
}