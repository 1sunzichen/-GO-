package main
import(
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	inputFile,inputError:=os.Open("input.txt")
	if inputError!=nil{
		fmt.Printf("An error occurred on opening the inputfile\n"+
	"Does")
		return 
	}
	defer inputFile.Close()
	inputReader:=bufio.NewReader(inputFile)
	for{
		inputString,readerError:=inputReader.ReadString('\n')
		fmt.Printf("the input was %s",inputString)
		if readerError== io.EOF{
			return 
		}
	}
}