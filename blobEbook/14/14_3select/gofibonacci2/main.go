package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//声明 长度
	tern := 25
	//声明 循环开始
	i := 0
	// 制作一个通道int类型的值
	c := make(chan int)
	// 开始计时
	start := time.Now()
	//启用 goruntime 调用方法 传入 长度,和 chan int 变量
	go fo(tern, c)
	//开启循环
	for {

		//接收返回的结果 , 如果成功
		if result, ok := <-c; ok {

			//打印
			fmt.Println("result", result)
			//循环变量增加
			i++
		} else {
			//接收不到数据
			//接收结束时间
			end := time.Now()
			//计算间隔
			delta := end.Sub(start)

			//打印
			fmt.Printf("时间%s\n", delta)
			//退出程序
			os.Exit(0)
		}
	}
}

//声明方法 接收  长度int 和 chan int
func fo(t int, c chan int) {

	//做一个循环
	for i := 0; i < t; i++ {

		//把计算过的值 给到 返回的变量
		c <- fofn(i)

	}
	//关闭通道
	close(c)
}

//斐波那契数列 接收int 返回 int
func fofn(i int) (r int) {
	//判断
	if i <= 1 {
		r = 1
	} else {
		r = fofn(i-1) + fofn(i-2)
	}
	//最后返回
	return r

}
