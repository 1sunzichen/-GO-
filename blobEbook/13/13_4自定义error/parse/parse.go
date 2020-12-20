package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	// Index int
	Word string
	// Err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
				fmt.Printf("Panicing %s\r \n", err)
			}
		}
	}()
	fields := strings.Fields(input)
	numbers = fields2numbers(fields)

	return
}
func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{field})
		}
		numbers = append(numbers, num)
	}
	return
}
