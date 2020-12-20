package main

import (
	"123go/blobEbook/13/13_4自定义error/parse"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 13.4",
		"2+2 =4",
		"1st class",
		"",
	}
	for i, ex := range examples {
		fmt.Printf("Parsing %v %q: \n", i, ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
