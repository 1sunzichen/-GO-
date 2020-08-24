package main
import(
	"fmt"
	"math"
)

func const2(){
	const filename string="abc.txt"
	const a,b=3,4
	var c2 int
	c2=int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c2)
}

func enums25(){
	//iota 自增值123
	const (
		cpp=iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp,javascript)
}
// func main(){
// 	const2();
// 	enums25()
// 	fmt.Println("Hello World")
// }
