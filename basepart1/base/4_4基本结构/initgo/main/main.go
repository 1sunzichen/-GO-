package main
//测试init方法 限于 main 函数之前执行
import(
	"fmt"
	trans "123go/basepart1/base/4_4基本结构/initgo"
)
var twoPi=2* trans.Pi
func main(){
  fmt.Printf("2*Pi=%g\n",twoPi)
}