package mock

import "fmt"
//函數 返回值的結構
type Retriever struct {
	Contents string
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
// 接受变成值类型
func (r Retriever) GetV(url string) string {
	return r.Contents
}

// 定义map类型 必须有返回值
func (r *Retriever)Post(url string,form map[string]string)string{
	r.Contents=form["contents"]
	return "ok"
}

func (r *Retriever)String() string{
	return fmt.Sprintf("Retriever:{Contents=%s}",r.Contents)
}