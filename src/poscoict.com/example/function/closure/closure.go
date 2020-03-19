package main

import "fmt"

func main() {
	a, b := 3, 5 // 바깥 함수의 변수

	// 익명 함수를 선언 및 정의
	anonyfunc := func(x int) int {

		// 바깥 함수의 변수와 익명 함수 내의 변수를 + 연산
		return a + b + x
	}

	r := anonyfunc(7) // 익명함수 사용

	fmt.Println(r) // 15
}
