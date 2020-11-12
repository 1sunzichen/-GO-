package main

import(
	"container/list"
	"unsafe"
	"fmt"
)
func main(){
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占中的字节数是 %d", n2, unsafe.Sizeof(n2))
	l:=list.New()
	l.PushFront("102")
	l.PushFront("101")
	l.PushBack("103")
// Front               Next
	for e:=l.Back();e!=nil;e=e.Prev() {
		fmt.Printf("%v\n",e.Value)
	}
}