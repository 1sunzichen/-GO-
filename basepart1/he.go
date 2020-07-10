package main

import "fmt"

var a, c, b = "1", 2, "999"

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDedu() {
	var a, s, d, f = 3, 4, true, "def"
	fmt.Println(a, s, d, f)
}

func main() {

	fmt.Println("helo world")
	variable()
	variableInit()
	variableTypeDedu()
}
