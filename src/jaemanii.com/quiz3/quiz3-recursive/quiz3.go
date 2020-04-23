package main

func f(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return f(n-1) + f(n-2)
}

func main() {

	print("output : ")
	println(f(5) % 1234567)
}
