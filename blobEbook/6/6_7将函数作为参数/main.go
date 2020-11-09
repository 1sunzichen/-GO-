package main
import "fmt"
func main()  {
	cb(4,Add)
}
func Add(a,b int){
	//decimal 十进制
	fmt.Printf("The sum of %d and %d is ：%d\n",a,b,a+b)

}
func  cb(y int,f func(y int,v int))  {
	f(y,2)
} 
