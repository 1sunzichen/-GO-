package main

import "io/ioutil"

func main() {
	const filename = "abc.txt"
	ioutil.ReadFile(filename)
}
