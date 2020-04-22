package main

import (
	"log"
	"os"
)

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal("error is -> " + err.Error())
	}
	println(f.Name())
}
