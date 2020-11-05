package main
import(
	"fmt"
)
func main(){
	i1:=5
	var intP *int //声明一个指针类型
	intP=&i1  
	//i1 %d 数字  &i1 %p取指针
	fmt.Printf("An integer: %d,it's %p 12 %p\n",i1,&i1,intP)
}