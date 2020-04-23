package main

func solution(n int) int {

	var num = make([]int, 0, n+1)
	num = append(num, 0)
	num = append(num, 1)

	for i := 2; i <= n; i++ {
		num = append(num, num[i-1]+num[i-2])
	}

	return num[n] % 1234567
}

func main() {

	print("output : ")
	println(solution(100000))
}
