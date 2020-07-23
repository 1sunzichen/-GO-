package fibonacci
func Fibonacci() func()int{
	a,s:=0,1
	return func() int {
		// 下一次啊 a 的位置 是 S   s的位置是和
		a,s=s,a+s
		//fmt.Println(a)
		return a
	}
}
