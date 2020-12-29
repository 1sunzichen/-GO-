package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
func ServeHTTP(w http.ResponseWriter,r *http.Request){
	//
	fmt.Fprint(w,r.URL.RawQuery)
	fmt.Fprint(w,r.URL.Host)
	fmt.Fprint(w,r.URL.Path)
	rawquery:=r.URL.RawQuery
	va,_:=url.ParseQuery(rawquery)
	va.Get("name")
	fmt.Fprintln(w,"*****key*****",va.Get("name"))
	for key:=range  r.Header{
		fmt.Fprintf(w,"%s:%s\n",key,r.Header[key])
	}
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
	//自定义启动一个http服务
	http.Handle("/userinfo",&define)
	http.HandleFunc("/getbody",GetHttpBody)

	http.HandleFunc("/formpost",FormPostForm)
	//
	http.HandleFunc("/MultipartForm",MultpartForm)
	http.ListenAndServe(":8080",nil)
	//url
	//uri
}

func GetHttpBody(w http.ResponseWriter,r *http.Request){
	len:=r.ContentLength
	body:=make([]byte,len)
	r.Body.Read(body)
	fmt.Fprintln(w,string(body))

}
/*
	form:url query,表单
	postform: 表单
*/
func FormPostForm(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
		fmt.Fprintln(w, "r.PostForm->",r.PostForm)
		fmt.Fprintln(w,"r.Form->",r.Form)
}
//
func MultpartForm(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(1024)
	feeders :=r.MultipartForm.File["file_to_upload"][0]
	f,err:= feeders.Open()
	if err!=nil{
		fmt.Fprintln(w,"fileHeader.Open()->",err.Error())
	}else{
		dataFormFile,err:=ioutil.ReadAll(f)
		if err ==nil{
			fmt.Fprintf(w,"%s\n",string(dataFormFile))
		}
	}
	fmt.Fprintln(w, "r.PostForm->",r.MultipartForm)

}