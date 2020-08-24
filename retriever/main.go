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

type Poster interface {
	Post(url string,form map[string]string) string
}
const url="http://plat.shangkehy.com/"
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster){
	poster.Post(url,
		map[string]string{
			"name":"ccmouse",
			"course":"golang",
		})
}
// 组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}
// 单一接口
type Retriever2 interface {
	GetV(url string) string
}

// 增加参数  增加增加的返回值类型
func session(s RetrieverPoster,a Retriever2)(string,string){

	s.Post(url,map[string]string {
		"contents":"尚客平台端",
	})

	return s.Get(url),a.GetV(url)
}

func main() {
	var r Retriever
	retr:=mock.Retriever{"this is a fake niupi com"}
	r = &retr
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
    // 传入第二个参数 为值类型
	fmt.Println(session(&retr,retr))
}

func inspect(r Retriever) {
	fmt.Printf(" > %T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
