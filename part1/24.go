package main

import (
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:

		panic(fmt.Sprintf("Wrong score: %d", score))

	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	}
	return g
}

//条件语句
// func main() {
// 	const filename = "24.txt"
// 	//方式1
// 	// contents, err := ioutil.ReadFile(filename)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	fmt.Printf("%s\n", contents)
// 	// }

// 	//方式2

// 	if contents, err := ioutil.ReadFile(filename); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%s\n", contents)
// 	}
// 	//方式2的contents 访问不到
// 	fmt.Println(
// 		grade(1),
// 		grade(40),
// 		grade(77),
// 		//grade(183),
// 	)
// }
