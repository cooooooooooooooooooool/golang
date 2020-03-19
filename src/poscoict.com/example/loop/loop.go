package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	sum = 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	println(sum)
}
