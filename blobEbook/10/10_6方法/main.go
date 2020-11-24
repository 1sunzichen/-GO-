package main
import (
	"fmt"
	"time"
	"strconv"
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
type List []int
func (l List) Len()int{return len(l)}
func (l *List) Append(val int){ *l=append(*l,val)}
func test1014(){
	//值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len:%d)",lst,lst.Len())
	//指针
	plst:=new(List)
	plst.Append(2)
	fmt.Printf("%v (len:%d)",plst,plst.Len())

}
type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
func test1015(){
	p:=new(Person)
	p.SetFirstName("Hello ")
	fmt.Println(p.FirstName())
}
type Card struct{
	wheelCount string
}
func(c *Card)numberOfWheels()string{

	return 	c.wheelCount
}
type Mercedes struct{
	Card

}
func (m *Mercedes)sayHiToMerkel(s string){
	m.wheelCount=s
}

//多重继承
type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func test1067(){
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}
type Base2 struct{
	id int
}
func (b *Base2)Id()int{
	return b.id
}
func (b *Base2)SetId(i int){
	b.id=i+10
}
type PersonBase struct{
	Base2
	FirstName string
	LastName string
}
type Employee struct{
	PersonBase
	salary string
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func test1011(){
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

func main(){
	// v:=employee{0.5}
	// fmt.Printf("%v\n",v.giveRaise())
	// m:=myTime{time.Now()}
	// fmt.Println("Full time now:",m.String())
	// fmt.Println("First 3  chars:",m.first3Chars())
	// test1013()
	// test1014()
	// m:=new(Mercedes)
	// m.sayHiToMerkel("kitty")
	// fmt.Println(m.numberOfWheels())
	// test1067()
	// E:=new(Employee)
	// E.SetId(8)
	// fmt.Println(E.Id())
	// test1011();
	test1020();
}


type Integer int
type Integer1 struct {n int}
func(p2 Integer1)get()int{
	return p2.n
}
func (i *Integer) String() string {
    return strconv.Itoa(int(*i))
}

func(p Integer)get()int{
	return int(p)
}
func f(i int){
	fmt.Println(i)
}
func test1020(){
	var v Integer
	v=9
	var v2 Integer1
	v2.n=10
	f(int(v))
	f(v2.n)
}