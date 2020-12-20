package main
func main(){
	//结构体 struct
	// 定义自定义类型
   //1.type
	type integer int
	var intVariables int
	var integerVaeiables integer
	intVariables=int(integerVaeiables)
	//struct 组合一定数量的字段完成 值类型
    type userinfo struct{
    	name string
    	age int
    	hoddy []string
    	moreinfo map[string]interface{}

	}
}