package main
import(
	"fmt"
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
func main(){
	fmt.Println(eval(3,4,"*"))
	q,r:=div(13,3)

	fmt.Println(q,r)
}