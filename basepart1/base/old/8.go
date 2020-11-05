package main
import(
	"fmt"
	"reflect"
	"runtime"
	"math"
)
func eval(a,b int,op string) (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_:=div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf(
			"unsupported operation: %s",op)
			//panic("unsupported operation:"+op)
	}
}
func div(a,b int)(q,r int){
	q=a/b
	r=a%b
	return 
}
func apply(op func(int,int) int,a,b int) int{
	p :=reflect.ValueOf(op).Pointer()
	opName :=runtime.FuncForPC(p).Name()
	fmt.Printf("CAll"+"(%d,%d)",opName,a,b)
	return op(a,b)
}
func pow(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}

func sum(numbers ...int) int{
	s :=0
	for i:= range numbers{
		s+=numbers[i]
	}
	return s
}
// func main(){
// 	fmt.Println(eval(3,4,"*"))
// 	q,r:=div(13,3)

// 	fmt.Println(q,r)
// 	fmt.Println(apply(pow,3,4))
// 	fmt.Println(sum(1,2,3,4))
// }