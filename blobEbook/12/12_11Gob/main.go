package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// 本实现给每一个数据类型都编译生成一个编解码程序，当单个编码器用于传递数据流时，会分期偿还编译的消耗，是效率最高的。
func main() {
	var newwork bytes.Buffer
	//NewEncoder返回一个将编码后数据写入w的*Encoder。
	enc := gob.NewEncoder(&newwork)
	//函数返回一个从r读取数据的*Decoder，如果r不满足io.ByteReader接口，则会包装r为bufio.Reader。

	dec := gob.NewDecoder(&newwork)
	//Encode方法将e编码后发送，并且会保证所有的类型信息都先发送。
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	err = enc.Encode(P{33, 43, 54, "Pythagoras222"})
	if err != nil {
		log.Fatal("encode error", err)
	}
	var q Q
	var qqq Q
	// Decode从输入流读取下一个之并将该值存入e。如果e是nil，将丢弃该值；否则e必须是可接收该值的类型的指针。如果输入结束，方法会返回io.EOF并且不修改e（指向的值）。
	err = dec.Decode(&q)
	err = dec.Decode(&qqq)
	if err != nil {
		log.Fatal("decode err", err)
	}
	fmt.Printf("%q:{%d,%d}\n", q.Name, q.X, q.Y)
	fmt.Printf("%q:{%d,%d}\n", qqq.Name, qqq.X, qqq.Y)
}
