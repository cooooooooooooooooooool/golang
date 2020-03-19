package main

import "fmt"

func printSlice(s []string) {

	fmt.Print("slice result : ")

	for i := 0; i < len(s); i++ {
		print(s[i])
	}

	println()
}

func main() {
	// 슬라이스 변수 선언하고 리터럴값으로 초기화
	var s = []string{"s", "l", "i", "c", "e"}
	printSlice(s)

	// 내장함수 make 로 길이 0 최대길이(capacity) 5 의 slice 생성
	var s2 = make([]string, 0, 5)
	printSlice(s2)

	// slice 변수에 문자열 추가
	s2 = append(s2, "s", "l", "i", "c", "e", "!")
	printSlice(s2)

	// 부분 slice
	s2 = s2[0:5]
	printSlice(s2)
}
