package main
import "fmt"


func printArray(arr [5]int){
	arr[0]=100;
	for i,v:=range arr{
		fmt.Println(i,v)
	}

}
func printArray2(arr *[5]int){
	arr[0]=100;
	for i,v:=range arr{
		fmt.Println(i,v)
	}

}
func main(){
	var arr1 [5] int
	arr2:=[3] int{1,3,5}
	arr3:=[...]int{2,4,6,8,120}
	// 4 行 5 列
	// var grid [4][5]int
	// fmt.Println(arr1,arr2,arr3)
	// fmt.Println(grid)
	printArray(arr1) //拷貝數組
	printArray2(&arr3)//指針改變
	// for i:=range arr3{
	// 	fmt.Println(arr3[i])
	// }
	//
	//數組是值類型   [1]int [2]int不同類型
	fmt.Println(arr1,arr2,arr3);// 值類型
}