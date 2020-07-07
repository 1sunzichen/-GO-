//package main
import (
	"fmt"
	"unicode/utf8"
)
func lengthOf15(s string) int{
	lastOccurred :=make(map[byte]int)
	start:=0
	maxLength:=0
	for i,ch:=range []byte(s){
		if lastI,ok:=lastOccurred[ch];
		ok&&lastI>=start{
			start=lastI+1
		}
		if lastOccurred[ch]>=start{
			start=lastOccurred[ch]+1
			
		}
		if i-start+1>maxLength{
			maxLength=i-start+1
		}
		lastOccurred[ch]=i
	}
	return maxLength;
}
func main15(){
	fmt.Println(
		lengthOf15("asbasnbass"))
	s:="Yes我愛93"
	fmt.Println(s)
	for _,b:=range []byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i,ch:=range s{//ch is a rune
		fmt.Printf("(%d %X) ",i,ch)
	}
	fmt.Println()
	fmt.Println("Rune count:",
	  utf8.RuneCountInString(s))
	bytes :=[]byte(s)

	for len(bytes)>0 {
		ch,size:= utf8.DecodeRune(bytes)
		bytes=bytes[size:]
		fmt.Printf("%c ",ch)
	}
}