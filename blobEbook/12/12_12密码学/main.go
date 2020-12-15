package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	//返回一个新的使用SHA1校验的hash.Hash接口。
	hasher := sha1.New()
	//WriteString函数将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法。
	io.WriteString(hasher, "test")
	b := []byte{}
	//返回数据data的SHA1校验和。
	fmt.Printf("Result:%x\n", hasher.Sum(b))
	fmt.Printf("Result:%d\n", hasher.Sum(b))
	// 重设hash为无数据输入的状态
	hasher.Reset()
	//a10 15 f 13 d 14 e   16-10 %x 16进制
	data := []byte("0123456789012345")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error:%v /%v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x %v\n", checksum, n)
}
