package main
import (
	"fmt"
	"os"
	"io/ioutil"
)
func data(name string)string{
	f, _ := os.OpenFile(name, os.O_RDONLY, 0)
	defer f.Close() // idiomatic Go code!
	contents, _ := ioutil.ReadAll(f)
	return string(contents)

}
func main(){
	fmt.Print(data("index.html"))
}