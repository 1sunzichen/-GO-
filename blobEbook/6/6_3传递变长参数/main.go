package main
import "fmt"
func main() {
	//x := min(1, 3, 2, 0)
	//fmt.Printf("The minimum is: %d\n", x)
	//slice := []int{7,9,3,5,1}
	//x = min(slice...)// 长参数传递...
	//fmt.Printf("The minimum in the slice is: %d", x)
	//deterfunc()
	//f()
	b()
}

//... 长参数接受
func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
func deterfunc(){
	fmt.Printf("In function1 at this top\n")
	defer funcation2()
}
func funcation2(){
	fmt.Printf("function2 call funcation")
}
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}