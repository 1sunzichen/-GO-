package main
import (
	"123go/learnGo/oop/model"
	"fmt"
)
func main(){
	boge:=model.NewUserInfo("架起",
		16,180,"北京有点",
		[]string{"跑步","泡妞"},nil)
	fmt.Printf("家奇的信息%v\n",boge)
}
