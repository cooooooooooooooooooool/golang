package main

import (
	"fmt"
	"strconv"
)

var result int = 1

func factorial(x int) int {
	if x > 0 {
		result *= x
		x--
		factorial(x)
	}

	return result
}

func main() {
	i := 12

	factorial(i)

	fmt.Println("결과값 : " + strconv.Itoa(result))
}
