package main
import (
	"fmt"
	"strings"
)
type struct1 struct{
	i1 int
	f1 float32
	str string
}
func main() {
	// ms:=new(struct1)
	// ms.i1=10
	// ms.f1=18.9
	// ms.str="Hello"
	// fmt.Printf("this int is:%d\n",ms.i1)
	// fmt.Printf("this int is:%f\n",ms.f1)
	// fmt.Printf("this int is:%s\n",ms.str)
	// fmt.Println(ms)//*取值  &取地址
	gogog()
}
type Person struct {
	name string
	age string
}
func upPerson(p Person){

	p.name = strings.ToUpper(p.name)
	p.age = strings.ToLower(p.age)
}
func upPerson2(p *Person){
	//指针改变
	p.name = strings.ToUpper(p.name)
	p.age = strings.ToLower(p.age)
}
func test1(a *int){
	*a=8
}
func gogog(){
	var p Person
	(&p).name ="xiaoming"
	p.age="20"
	upPerson2(&p)
	fmt.Printf("p int is:%s %s\n",p.name,p.age)

	p2:=new(Person)
	(*p2).name ="xiaoming2"
	p2.age="202"
	upPerson(*p2)
	fmt.Printf("p2 int is:%s %s\n",p2.name,p2.age)
	q:=9
	
	test1(&q)
	fmt.Printf("%d\n",q)
}