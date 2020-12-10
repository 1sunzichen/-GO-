package main
import(
	"reflect"
	"fmt"
)
type MyInt int

func main(){
	// var m MyInt = 5
	// v1 := reflect.ValueOf(m)
	// fmt.Printf("%v",v1.Kind())
	// var x float64 = 3.4
	// fmt.Println("type:", reflect.TypeOf(x))
	// v := reflect.ValueOf(x)
	// fmt.Println("value:", v)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind:", v.Kind())
	// fmt.Println("value:", v.Float())
	// fmt.Println(v.Interface())
	// fmt.Printf("value is %5.2e\n", v.Interface())
	// y := v.Interface().(float64)
	// fmt.Println(y)
	example();
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// v.SetFloat(9.0)
}	

func example(){
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}