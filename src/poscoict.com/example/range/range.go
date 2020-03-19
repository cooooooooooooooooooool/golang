package main

func main() {
	names := []string{"코로나", "감기", "바이러스"}

	for index, name := range names {
		println(index, name)
	}
}
