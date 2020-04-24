package main

import "os"

func main() {
	f, err := os.Open("C:\\temp\\PV3.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	n, err := f.Read(bytes)
	println(n)
	println(string(bytes))
}
