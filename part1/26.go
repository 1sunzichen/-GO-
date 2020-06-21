package main
import (
	"io/ioutil"
	"fmt"
)
func grade(score int) string{
	g:=""
	switch{
	case score<60:
		g="F"
	case score<80:
		g="C"
	case score<90:
		g="B"
	case score<=100:
		g="A"
	}
	return g
}

// func main(){
// 	const filename="abc.txt";
// 	//1
// 	//contents,err:=ioutil.ReadFile(filename);
// 	// if err!=nil{
// 	// 	fmt.Println(err)
// 	// }else{
// 	// 	fmt.Printf("%s\n",contents)
// 	// }
// 	if contents,err:=ioutil.ReadFile(filename);err==nil{
// 		fmt.Println(string(contents));
// 	}else{
// 		fmt.Println("cannot print file contents:",err)
// 	}

// 	fmt.Println(
// 		grade(89),
// 	)
// }

