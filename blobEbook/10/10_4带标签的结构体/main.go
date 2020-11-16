package main


import (
	"fmt"
	"reflect"
)

type TagType struct{
	field1 bool   "good"
	field2 string   "name"
	filed3 int      "mush"
}

func main(){
	tt:=TagType{true,"Barak Obama",1}
	for i:=0;i<3;i++{
		refTag(tt,i)
	}
}

func refTag(tt TagType,ix int){
	ttType:=reflect.TypeOf(tt)
	ixField:=ttType.Field(ix)
	fmt.Printf("%v\n",ixField.Tag)
}