package main

import (
	"123go/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	// fmt.Println(q)
	// fmt.Println(q.Pop())
	// fmt.Println(q.IsEmpty())
	// fmt.Println(q)
	q.Pop()

	fmt.Println(q)
	// q.Push("abc")
	// fmt.Println(q.Pop())
}

// go install 命令的功能和前面一节《go build命令》中介绍的 go build 命令类似，附加参数绝大多数都可以与 go build 通用。
// go install 只是将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下。
// 这个命令在内部实际上分成了两步操作：第一步是生成结果文件（可执行文件或者 .a 包），
// 第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。
