package main

import (
	"123go/retriever/mock"
	"123go/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://plat.shangkehy.com/")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake niupi com"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//类型， 值
	fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r))
	inspect(r)
	// Type assertion
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
