package main
import (
	"fmt"
	"strconv"
	
)
type TwoInts struct{
	b int
	a int
}
type Day int
const (
	Mo Day=iota
	Tu
	We
	Th
	Fr
	Sa
	Su
)
var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func(ddd Day)String()string{
	// switch D {
	// case condition:
		
	// }
	return dayName[ddd]
}
type Arrs [4]int
func(aaa *Arrs)Push2(v int){
	for i :=len(aaa)-1; i > 0; i-- {
		
		if aaa[i]==0{
			aaa[i]=v
			break
		}
	}

}
func(aaa *Arrs)Pop(v int){
	for _,aa:=range aaa{
		if aa==0{
			aa=v
			break
		}
	}
}
func(aaa Arrs)String()string{
	str:=""
	for _,ss:=range aaa{
		str+=strconv.Itoa(ss)+";"
	}
	
	return "数据为:"+str
}
func test1016(){
	var v Arrs
	v.Push2(4)
	fmt.Println(v)
}

func main(){
//    two1:=new(TwoInts)
//    two1.a=12
//    two1.b=2
//    fmt.Printf("two1 is: %T\n", two1)
// 	fmt.Printf("two1 is: %#v\n", two1)
//    var ce Celsius=9.9
//    fmt.Print(ce)
	
//    day:=Mo
//    var th Day=6
//    fmt.Println(day,th)
test1016()
	
}

func (tn *TwoInts) String()string{
	return "("+strconv.Itoa(tn.a)+"/"+strconv.Itoa(tn.b)+")"
}

type Celsius float64
func (ce Celsius) String()string{
	a:=strconv.FormatFloat(float64(ce),'f',2,64)

	b:=strconv.FormatFloat(float64(ce-32)* 5 / 9,'f',2,64)

	return  ""+a+","+b+"℃"
}
// func (ce *Celsius) String()string{
// 	return "("+strconv.FormatFloat(ce,"e",e,32)+","+(5/9*(ce-32))+"℃)"
// }