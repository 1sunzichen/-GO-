package main
import(
	"fmt"
	"strings"
	"time"
)
func main(){
	// f()
	// f2()
	start := time.Now()

	f := fibonacci()

	for i :=0; i < 41; i++ {
		fmt.Println(f())
	}
	// ooo:=MakeAddSuffix(".000")
	// fmt.Println(ooo("file.000"))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("斐波那契数列花费时长为 %s\n", delta)
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		// fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
func f2() (ret int) {
	defer func() {
		ret++
		fmt.Println(ret,"ret")
	}()
	return 1
}


func fibonacci() func() int {
	// i-2
	a := 1
	// i
    b := 0
    return func() int {

		//  a=> i-1   b =>i  a+b => i-1 + i 
		//     0         1, 1   
		fmt.Print(a," ") // i-2
 		a, b = b, a + b
		fmt.Print(a," ") // i-1 
        return b         // i
    }
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}