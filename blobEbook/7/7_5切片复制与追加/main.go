package main
import "fmt"
func main(){
	// slForm:=[]int{1,2,3}
	// slTo:=make([]int,2)
	// //复制里面的值
	// n:=copy(slTo,slForm)
	// fmt.Println(slTo,n)
	// sl3 := []int{1, 2, 3}
	// sl3 = append(sl3, 4, 5, 6)
	// fmt.Println(sl3)
	//
	// bytes:=make([]byte,2,2)
	// data:=[]byte{12,255,122,0,0,0}//0~255
	// fmt.Println(AppendByte(bytes,data...))
	//
	//fmt.Println(lenf([]int{1,2},4))
	//
//	fmt.Println(test710([]int{1,2,3,1,2,3,2,3,2,3},test710fn))
	arr2:=[6]int{0,1,2,3,4,5}
	
	fmt.Println(InsertStringSlice711(arr2[0:5],[]int{9,8,7},5))
	fmt.Println(arr2)
}
//b 0
func InsertStringSlice711(a []int,b []int,INDEX int)([]int,[]int){
	lens:=len(a)+len(b)
	res:=a
	a[0]=999
	if INDEX>len(res){
		return res,a
	}
	newSlice:=make([]int,lens)
	copy(newSlice,res)
	res=newSlice

    copy(res[len(b)+INDEX-1:],res[INDEX-1:])
	copy(res[INDEX-1:len(b)+INDEX-1],b)
	// res:=a
	return res,a
	
	
}

func test710fn(n int)bool{
	if n%2==0{
		return true
	}else{
		return false
	}
}
func test710(s []int,fn func(int) bool)[]int{
	 var res []int 
	
     for _,v:=range s{
		 if fn(v) {
			res=append(res,v)
		 }
	 }
	 return res
}

func lenf(s []int,factor int)([]int,int){
	
	nlen:=len(s)*factor
	newSlice:=make([]int,nlen)
	slice:=newSlice
	copy(slice[:len(s)],s)
	return slice,len(slice)
}

func AppendByte(slice []byte,data ...byte)[]byte{
	m:=len(slice)
	n:=m+len(data)
	if n>cap(slice){
		newSlice:=make([]byte,(n))
		copy(newSlice,slice)
		slice=newSlice
	}
	slice=slice[0:n]
	copy(slice[m:n],data)
	return slice
}