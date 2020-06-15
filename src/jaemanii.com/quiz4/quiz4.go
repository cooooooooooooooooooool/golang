package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var result int = 1

func solution(arg string) string {

	size := utf8.RuneCountInString(arg)

	str := strings.Split(arg, "")

	var answer string = string(str[size/2])

	if size%2 == 0 {
		start := size/2 - 1
		answer = string(str[start]) + string(str[start+1])
	}

	return answer
}

func main() {
	input := "1234567890"
	fmt.Println("결과값 : " + solution(input))
}
