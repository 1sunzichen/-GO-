package main
import (
	"fmt"
	"time"
)
type employee struct{
	salary float32

}
//方法：  现有结构体 后有方法
// 
func (th *employee) giveRaise()float32{
	return th.salary+1.0
}
//可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。
//或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。

type myTime struct{
	time.Time
}

func (t myTime) first3Chars() string{
	return t.Time.String()[0:3]
}
type B struct{
	thing int
}
func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func test1013(){
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	// b2 := new(B) // b2是指针
	// b2.change()
	// fmt.Println(b2.write())
}

func main(){
	// v:=employee{0.5}
	// fmt.Printf("%v\n",v.giveRaise())
	// m:=myTime{time.Now()}
	// fmt.Println("Full time now:",m.String())
	// fmt.Println("First 3  chars:",m.first3Chars())
	test1013()
}




