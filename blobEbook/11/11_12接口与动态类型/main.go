package main

import "fmt"

type Shaper interface{
	Area() float32
}

type TopologicalGenus interface{
	Rank() int
}

type Square struct{
	side float32
}

func (sq *Square)Area() float32{
	return sq.side * sq.side
}

func (sq *Square)Rank()int{
	return 1
}

type Rectangle struct{
	length,width float32
}

func (r Rectangle) Area()float32{
	return r.length * r.width
}

func (r Rectangle) Rank()int{
	return 2
}

func main(){
	r:=Rectangle{5,3}
	q:=&Square{8}
	shapes:=[]Shaper{r,q}
	fmt.Println("area")
	for n,_:=range shapes{
		fmt.Println("shapes 详情",shapes[n])
		fmt.Println("shapes的面积方法",shapes[n].Area())
	}
	topgen:=[]TopologicalGenus{r,q}
	for n,_:=range topgen{
		fmt.Println("topgen详情",topgen[n])
		fmt.Println("topgen返回常值的方法",topgen[n].Rank())
	}

}

