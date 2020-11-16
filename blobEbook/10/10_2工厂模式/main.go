package main
type Foo map[string]string
type Bar struct{
	thingOne string
	thingTwo int
}
// 现在为止我们已经见到了可以使用 make() 的三种类型中的其中两个：

// slices  /  maps / channels（见第 14 章）
func main(){
	//ok
	y:=new(Bar)
	(*y).thingOne="hello"
	(*y).thingTwo=1

	// NOT OK
    // z := make(Bar) // 编译错误：cannot make type Bar
    // (*z).thingOne = "hello"
    // (*z).thingTwo = 1

    // OK
    x := make(Foo)
    x["x"] = "goodbye"
    x["y"] = "world"

    // NOT OK
    // u := new(Foo)
    // (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
    // (*u)["y"] = "world"
}