package main

import "fmt"

func printArrayCut(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	//    arr:=[...]int{0,1,6,3,2,6,7}
	//    // 索引 3  -索引4  左包右閉
	//    	var arr132 [5] int
	// // slice 可以向後擴展 不可以向前擴展
	// 	arr332:=[...]int{2,4,6,8,120}
	//    s:=arr[2:4]//切片   引用傳遞
	//    s1:=s[1:4]
	// 	fmt.Println(s)
	// 	printArrayCut(arr132[:])
	// 	printArrayCut(arr332[:])
	// 	fmt.Println(arr332,arr132,s1)
	// 	s3:=append(s1,8);
	// 	s4:=append(s3,9);
	//    // s3  s4 改變視圖
	//    // 也不會修改 原來的長度
	// 	fmt.Println("arr=",arr,s3,s4,append(s4,10))
	//    var s32 []int
	//    for i:=0;i<100;i++{
	// 	printSlice(s32)
	// 	  s32=append(s32,i+1)
	//    }
	//    fmt.Println(s32);
	s321 := []int{2, 4, 6, 8}
	//   printSlice(s321)
	s322 := make([]int, 4)
	// s323:=make([] int,10,32)
	// printSlice(s322)
	// printSlice(s323)

	copy(s322, s321)

	printSlice(s322)
	//刪除 一個元素  包含左面 不含右面
	s322 = append(s322[:3], s322[4:]...)
	printSlice(s322)

	front := s322[0] //頭部刪除
	s322 = s322[1:]
	printSlice(s322)
	tail := s322[len(s322)-1]
	s322 = s322[:len(s322)-1] //尾部刪除
	printSlice(s322)
	fmt.Print(front, tail)
}
