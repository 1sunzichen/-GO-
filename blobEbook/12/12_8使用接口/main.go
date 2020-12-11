package main
import (
	"bufio"
	"fmt"
	"os"
	"io"
)
func main(){
	// fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// // buffered: os.Stdout implements io.Writer
	// buf := bufio.NewWriter(os.Stdout)
	// // and now so does buf.
	// fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	// buf.Flush()
	example()
}

func example(){
	inputFile,_:=os.Open("goprogram")
	outputFile,_:=os.OpenFile("goprogramT",os.O_WRONLY|os.O_CREATE,0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	var outputString string
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		if len(inputString) < 3 {
			outputString = "\r\n"
		} else if len(inputString) < 5 {
			outputString = string([]byte(inputString)[2:len(inputString)]) + "\r\n"
		} else {
        		outputString = string([]byte(inputString)[2:5]) + "\r\n"
		}
		// fmt.Print(outputString)
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	outputWriter.Flush()
	fmt.Println("Conversion done")
}