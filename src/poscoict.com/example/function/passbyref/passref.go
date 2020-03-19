package main

func printMsg(msg *string) {
	println(*msg)
	*msg = "Hi~ Hi~" // 넘겨받은 변수값을 변경
}

func main() {
	src := "Hi~"

	// Pass By Reference, 변수의 주소값을 파라미터로 전달
	printMsg(&src)

	// src 변수값의 변경 여부 확인
	println(src)
}
