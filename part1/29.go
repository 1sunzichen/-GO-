package main
import (
	"fmt"
)
// * 指針
func swap(a,b *int){
	*b,*a=*a,*b
}
// 
func swap2(a,b int)(int,int){
	return b,a
}
// func  main()  {
// 	a,b:=3,9;
// 	c,d:=3,9;
// 	//取地址
// 	swap(&a,&b);
// 	c,d=swap2(c,d)
// 	fmt.Println(c,d)
// }