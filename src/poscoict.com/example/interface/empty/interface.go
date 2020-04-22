package main

import "fmt"

func main() {
	var x interface{}
	x = "Tom"
	x = 1

	printIt(x)
}

func printIt(v interface{}) {
	fmt.Println(v)
}
