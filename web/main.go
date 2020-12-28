package main

import (
	"fmt"
	"net/http"
)
func ServeHTTP(w http.ResponseWriter,r *http.Request){
	//
	fmt.Fprint(w,r.URL.RawQuery)
	fmt.Fprint(w,r.URL.Host)
	fmt.Fprint(w,r.URL.Path)
	fmt.Fprintln(w,"通过handleFunc启动一个http")
}
type DefineServerMux struct { }
func (dsm *DefineServerMux)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"通过自定义类型启动一个http服务")
}
func main(){
 	//注册一个处理器函数
	http.HandleFunc("/",ServeHTTP)
	define:=DefineServerMux{}
	http.Handle("/userinfo",&define)
	http.ListenAndServe(":8080",nil)
	//url
	//uri
}
