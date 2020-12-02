package main
import "fmt"
type List []int
func (l List)Len() int{
	return len(l)
}
func (l *List) Append(val int){
	*l=append(*l,val)
}
type Appender interface{
	Append(int)
}

func CountInto(a Appender,start,end int){
	for i:=start;i<=end;i++{
		a.Append(i)
	}
}
type Lener interface{
	Len() int
}

func LongEnough(l Lener)bool{
	fmt.Print(l.Len())
	return l.Len()+43>42
}

// 在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个 Appender，
// 而它的方法 Append 只定义在指针上。 在 lst 上调用 LongEnough 是可以的，因为 Len 定义在值上。
// 在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender，
// 并且它的方法 Append 定义在指针上。 在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
func main(){
	var lst List
	// CountInto(lst, 1, 10)
	if LongEnough(lst){
		fmt.Printf("lst is long enough\n")
	}

	plst:=new(List)
	CountInto(plst,1,10)
	if LongEnough(plst){
		fmt.Printf("= plst is long enough\n")
	}
}