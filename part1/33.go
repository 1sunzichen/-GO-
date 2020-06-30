package main
import "fmt"

// func isok(p map[string],ok,mc map[string]){
// 	if p,ok:=mc;ok{
// 		fmt.Println(p)
// 	}else{
// 		fmt.Println("dont exist")
// 	}
// }
func main(){
	m:=map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	m2:=make(map[string]int)//empty
	var m3 map[string]int//nil
	fmt.Println(m,m2,m3)

	for k,v:=range m{
		//無序
		fmt.Println(k,v);
	}
	coursename:=m["course"]
	causename:=m["cause"]
	fmt.Println(causename)
	fmt.Println(coursename)
	// isok(causename,okc,m["cause"])
		if causename,ok:=m["cause"]; ok{
		fmt.Println(causename)
	}else{
		fmt.Println("dont exist")
	}
	
   // 刪除元素
   name:=m["name"]
   delete(m,"name")
   fmt.Println(name)

}