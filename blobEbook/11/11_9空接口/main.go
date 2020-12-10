package main
import "fmt"
type Any interface{}



var i = 5
var str="Abc"
type Person struct{
	name string
	age int
}
func main(){
	// var val Any
	// val = 5
	// fmt.Printf("val has the value: %v\n", val)
	// val = str
	// fmt.Printf("val has the value: %v\n", val)
	// pers1 := new(Person)
	// pers1.name = "Rob Pike"
	// pers1.age = 55
	// val = pers1
	// fmt.Printf("val has the value: %v\n", val)
	// switch t := val.(type) {
	// case int:
	// 	fmt.Printf("Type int %T\n", t)
	// case string:
	// 	fmt.Printf("Type string %T\n", t)
	// case bool:
	// 	fmt.Printf("Type boolean %T\n", t)
	// case *Person:
	// 	fmt.Printf("Type pointer to Person %T\n", t)
	// default:
	// 	fmt.Printf("Unexpected type %T", t)
	// }
	 example()
}
func FuncReturnSlice()[]int{
	return []int{2,3,4}
}
var dataSlice []int = FuncReturnSlice()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
func example(){
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	fmt.Print(interfaceSlice,"interfaceSlice")
}