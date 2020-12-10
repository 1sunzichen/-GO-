package main
import "fmt"
type Stringer interface{
	String() string
}
func main(){
	var v Stringer
	if sv,ok:=v.(Stringer); ok{
		fmt.Printf("v implements String():%s\n ",sv.String())
	}
}