package main

import (
	structPack	"123go/blobEbook/10/10_3自定义包"
	"fmt"
)
func main(){
	struct1:=new(structPack.Expstruct)
	struct1.Mi1=10
	struct1.Mf1=19.2
	fmt.Printf("Mi1= %d\n",struct1.Mi1)
	fmt.Printf("Mf1= %f\n",struct1.Mf1)
}