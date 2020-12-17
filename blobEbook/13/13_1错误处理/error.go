package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var errNotFound error = errors.New("Not found error12344566")

func main() {
	c()
	// fmt.Printf("error: %v", errNotFound)
}
func c() {

	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err := fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
		fmt.Print(err)
	}
}
