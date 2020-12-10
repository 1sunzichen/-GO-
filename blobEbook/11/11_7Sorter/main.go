package main
import "fmt"
type Sorter interface{
	Len() int
	Less(i,j int)bool
	Swap(i, j int)
}
func Sort(data Sorter){
	//
	for pass := 1; pass < data.Len(); pass++ {
		// 1-0，2-1，3-2 ,4-3,	4   把最大的放到右面
		// 1-0, 2-1  3-2 		3   把最大的放到右面
		// 1-0, 2-1            	2   把最大的放到右面
		// 1-0                  1   把最大的放到右面
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1,i){
				data.Swap(i,i+1)
			}
		}
	}
}
func IsSorted(data Sorter) bool{
	n:=data.Len()
	for i:=n-1;i>0;i--{
		if data.Less(i,i-1){
			return false
		}
	}
	return true
}
type FloatArray []float64
func(p FloatArray) Len() int{return len(p)}
func(p FloatArray) Less(i,j int) bool{return p[i]<p[j]}
func(p FloatArray) Swap(i,j int) {p[i],p[j]=p[j],p[i]}
type IntArray []int

func(p IntArray) Len() int{return len(p)}
func(p IntArray) Less(i,j int) bool{return p[i]<p[j]}
func(p IntArray) Swap(i,j int) {p[i],p[j]=p[j],p[i]}

type StringArray []string

func(p StringArray)Len() int{return len(p)}
func(p StringArray)Less(i,j int) bool{return p[i]<p[j]}
func(p StringArray)Swap(i,j int) {p[i],p[j]=p[j],p[i]}

func Sortints(a []int){Sort(IntArray(a))}
func SortStrings(a []string){Sort(StringArray(a))}
func SortFloat(a FloatArray){Sort(FloatArray(a))}

func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
func FloatsAreSorted(a FloatArray) bool       { return IsSorted(FloatArray(a)) }
func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }

func main(){
	a:=FloatArray{12.2,13.12,1.1}
	SortFloat(a)
	fmt.Print(a)
	// example();
}

type Interface interface{
	Len() int
	Less(i,j int)bool
	Swap(i,j int)
}
type Triangle struct{
	p float32
	q float32
}

type AreaInterface interface{
	Area() float32
}

func (t Triangle)Area()float32{
	return 0.5 * t.p * t.q
}

type Square struct{
	p int
	q int
}

type PeriInterface interface{
	Perimeter() int
}
func (s Square)Perimeter()int{
	return (s.q+s.p)*2
}

func example(){
	// var a AreaInterface
	// a=Triangle{10.2,11.9}
	// fmt.Print(a.Area())
	b:=Square{9,10}
	c:=[]PeriInterface{b}
	fmt.Print(c[0].Perimeter())

}
