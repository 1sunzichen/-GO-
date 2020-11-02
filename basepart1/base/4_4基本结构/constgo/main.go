package main
import(
	"fmt"
)
func demo0(){

	const(
		Unknown=0
		Female=1
		Male=0
	)
	fmt.Println(Unknown,Female,Male)
}

func demo()  {
	
	const (
			a=iota
			b=iota
			c=iota
	)
	fmt.Println(a,b,c)
}
func demo2(){

	const (
		a = iota  // a = 0
		b         // b = 1
		c         // c = 2
		d = 5     // d = 5   
		e         // e = 5
	)
	fmt.Println(a,b,c,d,e)
}

func demo3(){
	const (
		Apple, Banana = iota + 1, iota + 2 // Apple=1 Banana=2
		Cherimoya, Durian                  // Cherimoya=2 Durian=3
		Elderberry, Fig                    // Elderberry=3, Fig=4
	
	)
	const (
		Open = 1 << iota  // 0001
		Close             // 0010
		Pending           // 0100
	)
	
	const (
		_           = iota             // 使用 _ 忽略不需要的 iota
		KB = 1 << (10 * iota)          // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)
	fmt.Println(Apple,Banana,Cherimoya,Durian,Elderberry,Fig,Open,Close,Pending)
	fmt.Println(KB,MB,GB,TB,PB,EB)
	fmt.Println(YB/EB,"%T/%T",ZB/EB)
}

func vardemo(){
	var a=15
	b:=false
	var i=7
	j:=i
	j=9
	fmt.Println(a,b,"qwe`",i,j)
}
func main(){
	// demo2();
	// demo3();
	vardemo()
}