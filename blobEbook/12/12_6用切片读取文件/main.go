package main
import(
	"flag"
	"fmt"
	"os"
)
func cat(f *os.File){
	const NBUF=2
	var buf [NBUF]byte
	//无限循环
   
	for{
		// fmt.Print(buf[:],"err")
		// f.Read(buf[:]);
		//
		// switch nr,err:=f.Read(buf[:]);true{
		switch nr,err:=f.Read(buf[:]);{
		case nr<0:
			fmt.Fprintf(os.Stderr,"cat:error reading:%s\n",err.Error())
			os.Exit(1)
		case nr==0:
			fmt.Print(nr,err,"err")
			return 
		case nr>0:
			if nw,ew:=os.Stdout.Write(buf[0:nr]);nw!=nr{
				fmt.Fprintf(os.Stderr,"cat: error reading: %s\n",ew.Error())
			}
		}
	}
}
func main(){
		flag.Parse()
		if flag.NArg()==0{
			cat(os.Stdin)
		}
		for i:=0;i<flag.NArg();i++{
			f,err:=os.Open(flag.Arg(i))
			if f==nil{
				fmt.Fprintf(os.Stderr,"cat: can't open %s:error%s\n",flag.Arg(i),err)
			}
			cat(f)
			defer f.Close()
		}
}