//打印出什么值
// 找作用域a的值 就可以了
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   a := "O"  // "G"  "O" "G"
   // a="O"  // "G" "O" "O" 修改之前的值
   print(a)
}


