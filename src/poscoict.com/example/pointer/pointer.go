package main

import (
	"fmt"
	"strconv"
)

func main() {
	var k int = 10
	var p = &k
	fmt.Println("포인터 p 의 주소에 해당하는 값 : " + strconv.Itoa(*p))
}
