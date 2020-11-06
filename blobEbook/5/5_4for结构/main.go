package main

import (
	"fmt"
)



func main(){
 
	// for1()
	// for2()
	// for3()
	 
	// stringfor()
	// for5test()
	//forrange1()
	fortest59()
}
func for3(){
	for i:=0; i<5; i++ {
		for j:=0; j<10; j++ {
			fmt.Printf("%d",j)
		}
	}
}
func stringfor(){
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c ", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c ", ix, str2[ix])
	}
}
func for1(){
	for i := 0; i < 5; i++ {
		fmt.Printf("this is this %d \n",i)
	 }
}
func for2(){
	N:=5
	for i, j := 0, N;  i < j;  i, j = i+1, j-1 {
	   fmt.Printf("this is this %d %d %v  \n",i,j,N)

	}
}

func for5test(){
	// for i:=1;i<16;i++{
	// 	fmt.Print("\n")
	// 	for j := 0; j < i; j++ {
	// 		fmt.Printf("%s","G")
	// 	}
	// }
	// i:=0
	// //goto重写
	// START:
	// fmt.Printf("the counter os %d\n",i)
	// i++
	// if i<15 {
	// 	goto START
	// }
	// for i:=0;i<=10;i++{
	// 	//%b 二进制 ^i 补码
	// 	fmt.Printf("the componnent of %b is %b\n",i,^i)
	// }
	// for v:=0;v<=100;v++{
	// 	//%b 二进制 ^i 补码
		
	// 	switch {
	// 	case v%15==0:
	// 		fmt.Printf("%s\n","FizzBuzz")
	// 	case v%3==0:
	// 		fmt.Printf("%s\n","Fizz")
	// 	case v%5==0:
	// 		fmt.Printf("%s\n","Buzz")
	// 	default:
	// 		fmt.Printf("%d\n",v)
	// 	}
	// }
	for v:=0;v<=10;v++{
		//%b 二进制 ^i 补码
		 fmt.Printf("\n")
		for i := 0; i < 20; i++ {
			fmt.Printf("%s","*")
		}
	}

}
func forrange1(){
	// str := "zd is a beautiful girl!"
	// // %c 字符
	// for pos,char :=range str{
	// 	fmt.Printf("index:%d,val:%c \n",pos,char)
	// }
	fmt.Println()
	str2 := "Chinese: 日本語"
	// fmt.Printf("The length of str2 is: %d\n", len(str2))
	// for pos, char := range str2 {
    // 	fmt.Printf("character %c starts at byte position %d\n", char, pos)
	// }
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
    	fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
func fortest59(){
	// for i := 0; i < 5; i++ {
	// 	var v int
	// 	fmt.Printf("%d ", v)
	// 	v = 5
	// } //没有判断条件 死循环
	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }//没有终止条件 循环
	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }
	// s := "" // , a ,aa,aaa,aaaa
	// for ; s != "aaaaa"; {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }
	// a,aa,aaa
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	s = i+1, j+1, s + "a" {
	fmt.Println("Value of i, j, s:", i, j, s)
}
}