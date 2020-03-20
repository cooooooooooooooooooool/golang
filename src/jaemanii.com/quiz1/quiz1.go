package main

import (
	"fmt"
	"strconv"
)

var result int = 0

func sum(in string) int {
	strArr := []rune(in)

	for i := 0; i < len(strArr); i++ {
		x, _ := strconv.Atoi(string(strArr[i]))
		result += x
		fmt.Println("result : " + strconv.Itoa(result))
	}

	return result
}

func main() {
	L := 13
	N := 17

	for ; L <= N; L++ {
		sum(strconv.Itoa(L))
	}

	fmt.Println("결과값 : " + strconv.Itoa(result))
}
