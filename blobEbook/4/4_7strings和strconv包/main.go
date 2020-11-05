package main

import(
	"fmt"
	"strings"
	"strconv"
)

func main(){
	// s:="prefixqqq"  %s string  %t bool %v int
	// //前缀和后缀 HasPrefix HasSuffix
	// fmt.Println(strings.HasPrefix(s,"prefix"))
	var str string = " This is an example of a string"
	// a:=strings.Fields(str)
	// b:=strings.Split(str,"a")
	val, err := strconv.Atoi(str)
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t%t\n%v%v\n%v%v \n%s %v\n %s \n%s %s\n %s %s %s  %v\n%v %v *********%v",
	strings.HasPrefix(str, "Th"),strings.Contains(str,"example"),
	strings.Index(str,"e"),strings.LastIndex(str,"rrr"),
	strings.IndexRune(str,rune('q')),strings.IndexRune(str,11),
	strings.Replace(str,"a","b",-1),strings.Count(str,"a"),
	strings.Repeat(str,2),
	strings.ToLower(str),strings.ToUpper(str),
	strings.TrimSpace(str),strings.Fields(str),strings.Split(str,"a"),strings.NewReader(str),
	strconv.Itoa(98),err,val,

	)


}