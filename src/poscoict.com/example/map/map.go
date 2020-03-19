package main

import "fmt"

func main() {
	// map 선언, 초기화는 make 함수를 사용
	var mymap = make(map[int]string)
	mymap[0] = "m"
	mymap[1] = "ap"

	// map 키 체크
	_, exist := mymap[2]
	if !exist {
		println("No map's key!")
	}

	for key, val := range mymap {
		fmt.Println(key, val)
	}

	println("--------------------")

	// key 0 을 삭제
	delete(mymap, 0)

	for key, val := range mymap {
		fmt.Println(key, val)
	}
}
