package flagdemo

import (
	"fmt"
	"flag"
)
func Flagdemos(){
	wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")
 
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var") // 对变量取址
 
    flag.Parse()
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr) // 对指针取值
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())

}