package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int
	num1 = 100
	fmt.Println("variable num1 : " + strconv.Itoa(num1))
	num2 := 200
	fmt.Println("variable num2 : " + strconv.Itoa(num2))
	const s1 string = "Hello"
	fmt.Println("variable s1 : " + s1)
	const s2 = "Hello2"
	fmt.Println("variable s2 : " + s2)
}
