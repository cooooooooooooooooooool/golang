package main

import (
	"fmt"
	"strconv"
)

// struct 정의
type user struct {
	id   string
	name string
	age  int
}

// toString 함수 정의
func toString(u user) string {
	return "id : " + u.id + ", name : " + u.name + ", age : " + strconv.Itoa(u.age)
}

func main() {
	// user 기본 객체 생성
	u1 := user{}

	// 필드값 설정
	u1.id = "test"
	u1.name = "Lee"
	u1.age = 10

	// 필드값을 초기화하여 user 객체 생성
	u2 := user{"test2", "Kim", 20}

	fmt.Println("user1 : " + toString(u1))
	fmt.Println("user2 : " + toString(u2))
}
