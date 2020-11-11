package main

import(
	"fmt"
)
const LIM=50

func main(){
//   var arr1 [7]int
// 	for i := 0; i < len(arr1); i++ {
// 	arr1[i]=i*2	
// 	}
// 	for i := 0; i < len(arr1); i++ {
// 		fmt.Printf("Array at index%d %d\n",i,arr1[i])
// 	}
	// slicefn()
	//
	// fb:=fibonacciArray(50)
	// fmt.Println(fb);
	// constarr()

	// arr:=multidimarray();
	// fmt.Println(arr,"arr")
}
func slicefn(){
	a:=[...]string{"a","b","c","d"}
	for i := range a{
		fmt.Println("Array item",i,"is",a[i])
	}
	//引用类型
	var arr1 = new([5]int)
    // 值类型
	var arr2 [5]int
	arr3:=  arr1
	arr4:=  &arr2
	arr5:=  *arr1
	arr6:=  arr2
	arr5[2] = 100
	arr6[2] = 100
	fmt.Println(arr1,arr2,arr3,arr4,arr5,arr6)
}
func fibonacciArray(count int) []int{
	farr:=make([]int,count)
	farr[0],farr[1]=1,1
	for i := 2; i < count; i++{
		farr[i]=farr[i-1]+farr[i-2]
	}		
	return farr
}

func constarr(){
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = []string{3: "Chris", 4: "Ron"}
	 //var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	 constest(arrKeyValue...)////切片被打散传入
}
//可以接受任意个string参数
func constest(arrKeyValue ...string){
	for i:=0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
const (
	WIDTH=10
	HEIGHT=10
)
type pixel int
var screen [WIDTH][HEIGHT]pixel
func multidimarray()[WIDTH][HEIGHT]pixel{ 
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			screen[j][i]=0
		}
	}
	return screen
}
func Sum(a *[3]float64)(sum float64){
	for _,v:=range a{
		sum+=v
	}
	return 
}
func init(){
	arrs:=[3]float64{3.2,5.7,2.4}
	x:=Sum(&arrs)
	fmt.Printf("%f",x)
}
					