package main

import (
	"log"
	"os"
)

var logger *log.Logger

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer fpLog.Close()

	logger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Print("Test")

	logger.Println("End of Program")

	logger.Fatal("exit!")
}
