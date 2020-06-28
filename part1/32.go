package main
import "fmt"


func printArrayCut(arr []int){
	arr[0]=100
	for i,v:=range arr{
		fmt.Println(i,v)
	}
}
func main(){
   arr:=[...]int{0,1,6,3,2,6,7}
   // 索引 3  -索引4  左包右閉
   	var arr132 [5] int
// slice 可以向後擴展 不可以向前擴展
	arr332:=[...]int{2,4,6,8,120}
   s:=arr[2:4]//切片   引用傳遞
   s1:=s[1:4]
	fmt.Println(s)
	printArrayCut(arr132[:])
	printArrayCut(arr332[:])
	fmt.Println(arr332,arr132,s1)
	s3:=append(s1,8);
	s4:=append(s3,9);
   // s3  s4 改變視圖
   // 也不會修改 原來的長度
	fmt.Println("arr=",arr,s3,s4,append(s4,10))

}